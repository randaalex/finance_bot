package alfamail

import (
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

type AlfaParser struct {
	accounts   *[]entities.Account
	categories *[]entities.Category
	logger     *logrus.Logger
}

func NewAlfaParser(logger *logrus.Logger, accounts *[]entities.Account, categories *[]entities.Category) *AlfaParser {
	return &AlfaParser{
		accounts:   accounts,
		categories: categories,
		logger:     logger,
	}
}

func (p *AlfaParser) Parse(mail string) *entities.Transaction {
	// TODO: return error

	//if isValid, _ := regexp.Match(cardNumberRegexp, []byte(mail)); !isValid {
	//	fmt.Printf("parse error: %v | %v\n", errors.New("invalid mail format"), mail)
	//	return &entities.Transaction{}
	//}

	mailRows := strings.Split(mail, "\n")

	account := p.getAccount(mail)
	amount, foreignAmount, foreignCurrencyCode := p.getAmounts(mail)
	transactionType := p.getType(mail)
	issuedAt := p.getTime(mail)

	transaction := entities.Transaction{
		IssuedAt:     issuedAt,
		Type:         transactionType,
		Description:  mailRows[7],
		Amount:       float32(amount),
		CurrencyCode: account.CurrencyCode,
		Recipient:    p.getRecipient(mail),
		Notes:        mail,
		//CategoryId:          fireflySplit.GetCategoryId(),
		//CategoryName:        fireflySplit.GetCategoryName(),
	}

	if foreignAmount != nil {
		transaction.ForeignAmount = float32(*foreignAmount)
		transaction.ForeignCurrencyCode = *foreignCurrencyCode
	}

	if transactionType == entities.TransactionSplitTypeWithdrawal {
		transaction.SourceId = account.Id
	} else {
		transaction.DestinationId = account.Id
	}

	return &transaction
}

func (p *AlfaParser) getType(value string) string {
	switch value {
	case "Оплата товаров/услуг", "Перевод (Списание)":
		return entities.TransactionSplitTypeWithdrawal
	case "Поступление успешно":
		return entities.TransactionSplitTypeDeposit
	}

	// TODO: error?
	return entities.TransactionSplitTypeWithdrawal
}

func (p *AlfaParser) getAccount(mail string) *entities.Account {
	const cardNumberRegexp = `(?m)^Карта \d.(\d{4})`

	re := regexp.MustCompile(cardNumberRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		panic("parse error")
	}

	for _, acc := range *p.accounts {
		if acc.AccountNumber == match[1] {
			return &acc
		}
	}

	panic("account not found")
}

func (p *AlfaParser) getTime(mail string) time.Time {
	const timeRegexp = `(?m)(\d{2}.\d{2}.\d{4} \d{2}:\d{2}:\d{2})`

	re := regexp.MustCompile(timeRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		panic("parse error")
	}

	parsedTime, err := time.ParseInLocation(timeFormat, match[1], entities.DefaultTZ())
	if err != nil {
		panic(err)
	}

	return parsedTime
}

func (p *AlfaParser) getAmounts(value string) (
	amount float64, foreignAmount *float64, foreignCurrency *string,
) {
	const amountRegexp = `(?m)^Сумма:(\d+\.?\d{0,2})\s(\w{3})(\s\((\d+\.?\d{0,2})\s(\w{3})\))?$`

	re := regexp.MustCompile(amountRegexp)
	match := re.FindStringSubmatch(value)

	if match == nil {
		panic("parse error")
	}

	amount = 0.0
	parenthesisAmount := 0.0
	currency := match[2]
	parenthesisCurrency := match[5]

	amount, _ = strconv.ParseFloat(match[1], 64)
	parenthesisAmount, _ = strconv.ParseFloat(match[4], 64)

	if parenthesisCurrency == "" {
		// transaction in native currency
		return amount, nil, nil
	} else {
		// transaction with exchange
		if currency != "USD" && currency != "EUR" && currency != "BYN" {
			fmt.Printf("unknown currency: %v\n", currency)

			panic("unknown currency")
		}

		return parenthesisAmount, &amount, &currency
	}
}

func (p *AlfaParser) parseBalance(mail string) float64 {
	const balanceRegexp = `(?m)^Остаток:(\d+\.?\d{0,2})\s(\w{3})$`

	re := regexp.MustCompile(balanceRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		return 0.0
	}

	balance, _ := strconv.ParseFloat(match[1], 64)

	return balance
}

func (p *AlfaParser) getRecipient(mail string) string {
	const detailsRegexp = `(?m)^(.*\/.*\/.*)$`

	re := regexp.MustCompile(detailsRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		panic("parse error")
	}

	return match[1]
}
