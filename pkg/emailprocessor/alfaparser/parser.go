package alfaparser

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/charmap"

	"github.com/randaalex/finance_bot/pkg/entities"
)

const (
	timeFormat = "02.01.2006 15:04:05"
)

var encoding = charmap.Windows1251

type Parser struct {
	accounts   *[]entities.Account
	categories *[]entities.Category
	logger     *logrus.Logger
}

func NewParser(logger *logrus.Logger, accounts *[]entities.Account, categories *[]entities.Category) *Parser {
	return &Parser{
		accounts:   accounts,
		categories: categories,
		logger:     logger,
	}
}

func (p *Parser) Parse(mail string) (*entities.Transaction, error) {
	if mail == "" {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "parse",
			Err:    errors.New("invalid mail error"),
		}
	}

	account, err := p.getAccount(mail)
	if err != nil {
		return nil, err
	}

	transactionType, err := p.getType(mail)
	if err != nil {
		return nil, err
	}

	if *transactionType == entities.TransactionSplitTypeTransfer {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "parse",
			Err:    errors.New("skip mail"),
		}
	}

	description, err := p.getDescription(*transactionType, mail)
	if err != nil {
		return nil, err
	}

	amount, foreignAmount, foreignCurrencyCode, err := p.getAmounts(mail)
	if err != nil {
		return nil, err
	}

	issuedAt, err := p.getTime(mail)
	if err != nil {
		return nil, err
	}

	recipient, err := p.getRecipient(*transactionType, mail)
	if err != nil {
		return nil, err
	}

	transaction := entities.Transaction{
		IssuedAt:     *issuedAt,
		Type:         *transactionType,
		Description:  *description,
		Amount:       *amount,
		CurrencyCode: account.CurrencyCode,
		Recipient:    *recipient,
		Notes:        mail,
	}

	if foreignAmount != nil {
		transaction.ForeignAmount = *foreignAmount
		transaction.ForeignCurrencyCode = *foreignCurrencyCode
	}

	if *transactionType == entities.TransactionSplitTypeWithdrawal {
		transaction.SourceId = account.Id
	} else {
		transaction.DestinationId = account.Id
	}

	return &transaction, nil
}

func (p *Parser) getType(mail string) (*string, error) {
	withdrawal := entities.TransactionSplitTypeWithdrawal
	deposit := entities.TransactionSplitTypeDeposit
	transfer := entities.TransactionSplitTypeTransfer

	mailRows := strings.Split(mail, "\n")
	if len(mailRows) < 2 {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "getType",
			Err:    errors.New("match transaction type (2) line error"),
		}
	}

	switch mailRows[2] {
	case "Оплата товаров/услуг", "Перевод (Списание)":
		return &withdrawal, nil
	case "Поступление", "Поступление успешно":
		return &deposit, nil
	case "Списание", "Выдача наличных":
		return &transfer, nil
	}

	return nil, ParseMailError{
		Mail:   mail,
		Action: "getType",
		Err:    fmt.Errorf("unknown transaction type '%s'", mailRows[2]),
	}
}

func (p *Parser) getDescription(transactionType, mail string) (*string, error) {
	if transactionType == entities.TransactionSplitTypeDeposit {
		s := "income"
		return &s, nil
	}

	mailRows := strings.Split(mail, "\n")

	if len(mailRows) < 8 {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "getDescription",
			Err:    errors.New("match description (7) line error"),
		}
	}

	description := mailRows[7]

	return &description, nil
}

func (p *Parser) getAccount(mail string) (*entities.Account, error) {
	const cardNumberRegexp = `(?m)^Карта \d.(\d{4})`

	re := regexp.MustCompile(cardNumberRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "getAccount",
			Err:    errors.New("match account error"),
		}
	}

	for _, acc := range *p.accounts {
		if acc.AccountNumber == match[1] {
			return &acc, nil
		}
	}

	return nil, ParseMailError{
		Mail:   mail,
		Action: "getAccount",
		Err:    fmt.Errorf("unknown account '%s'", match[1]),
	}
}

func (p *Parser) getTime(mail string) (*time.Time, error) {
	const timeRegexp = `(?m)На время:(\d{2}:\d{2}:\d{2})`
	const dateRegexp = `(?m)(\d{2}\.\d{2}\.\d{4})`

	reTime := regexp.MustCompile(timeRegexp)
	matchTime := reTime.FindStringSubmatch(mail)

	if matchTime == nil {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "getTime",
			Err:    errors.New("match time error"),
		}
	}

	reDate := regexp.MustCompile(dateRegexp)
	matchDate := reDate.FindStringSubmatch(mail)

	if matchDate == nil {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "getTime",
			Err:    errors.New("match date error"),
		}
	}

	dateTime := matchDate[1] + " " + matchTime[1]

	parsedDateTime, err := time.ParseInLocation(timeFormat, dateTime, entities.DefaultTZ())
	if err != nil {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "getTime",
			Err:    errors.New("parse datetime error"),
		}
	}

	return &parsedDateTime, nil
}

// return: amount, foreignAmount, foreignCurrency, error
func (p *Parser) getAmounts(mail string) (*float64, *float64, *string, error) {
	const amountRegexp = `(?m)^Сумма:(\d+\.?\d{0,2})\s(\w{3})(\s\((\d+\.?\d{0,2})\s(\w{3})\))?$`

	re := regexp.MustCompile(amountRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		return nil, nil, nil, ParseMailError{
			Mail:   mail,
			Action: "getAmounts",
			Err:    errors.New("match amounts error"),
		}
	}

	amount := 0.0
	currency := match[2]

	amount, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return nil, nil, nil, ParseMailError{
			Mail: mail,
			Line: match[1],
			Action: "ParseFloat",
			Err:  err,
		}
	}

	parenthesisCurrency := match[5]

	if parenthesisCurrency == "" {
		// transaction in native currency

		return &amount, nil, nil, nil
	} else {
		// transaction with exchange
		parenthesisAmount, err := strconv.ParseFloat(match[4], 64)
		if err != nil {
			return nil, nil, nil, ParseMailError{
				Mail: mail,
				Line: match[4],
				Action: "ParseFloat",
				Err:  err,
			}
		}

		if currency != "USD" && currency != "EUR" && currency != "BYN" {
			// TODO: logger, sentry
			return &parenthesisAmount, nil, nil, nil
		}

		return &parenthesisAmount, &amount, &currency, nil
	}
}

func (p *Parser) parseBalance(mail string) (*float64, error) {
	const balanceRegexp = `(?m)^Остаток:(\d+\.?\d{0,2})\s(\w{3})$`

	re := regexp.MustCompile(balanceRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "parseBalance",
			Err:    errors.New("match balance error"),
		}
	}

	balance, _ := strconv.ParseFloat(match[1], 64)

	return &balance, nil
}

func (p *Parser) getRecipient(transactionType, mail string) (*string, error) {
	const recipientRegexp = `(?m)^(.*\/.*\/.*)$`

	if transactionType == entities.TransactionSplitTypeDeposit {
		s := ""
		return &s, nil
	}

	re := regexp.MustCompile(recipientRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		return nil, ParseMailError{
			Mail:   mail,
			Action: "getRecipient",
			Err:    errors.New("match recipient error"),
		}
	}

	recipient := match[1]

	return &recipient, nil
}
