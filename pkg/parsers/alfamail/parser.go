package alfamail

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
	timeFormat       = "02.01.2006 15:04:05"
	timeRegexp       = `(?m)(\d{2}.\d{2}.\d{4} \d{2}:\d{2}:\d{2})`
	amountRegexp     = `(?m)^Сумма:(\d+\.?\d{0,2})\s(\w{3})(\s\((\d+\.?\d{0,2})\s(\w{3})\))?$`
	cardNumberRegexp = `(?m)^Карта \d.(\d{4})`
	balanceRegexp    = `(?m)^Остаток:(\d+\.?\d{0,2})\s(\w{3})$`
	detailsRegexp    = `(?m)^(.*\/.*\/.*)$`
)

var encoding = charmap.Windows1251

type AlfaParser struct {
	logger *logrus.Logger
}

func NewAlfaParser(logger *logrus.Logger) *AlfaParser {
	return &AlfaParser{
		logger: logger,
	}
}

func (a *AlfaParser) Parse(mail string) *[]entities.ParsedTransaction {
	var parsedTransactions []entities.ParsedTransaction

	if isValid, _ := regexp.Match(cardNumberRegexp, []byte(mail)); !isValid {
		fmt.Printf("parse error: %v | %v\n", errors.New("invalid mail format"), mail)
		return &parsedTransactions
	}

	mailRows := strings.Split(mail, "\n")

	parsedCardNumber := a.parseCardNumber(mail)
	parsedTime := a.parseTime(mail)
	parsedType := a.parseTransactionType(mailRows[2])
	parsedAmount, parsedCurrencyCode, parsedForeignAmount, parsedForeignCurrencyCode := a.parseAmountAndCurrency(mail)
	parsedBalance := a.parseBalance(mail)

	parsedTransactions = append(parsedTransactions, entities.ParsedTransaction{
		CardNumber:          parsedCardNumber,
		Time:                parsedTime,
		Type:                parsedType,
		Amount:              float32(parsedAmount),
		CurrencyCode:        parsedCurrencyCode,
		ForeignAmount:       float32(parsedForeignAmount),
		ForeignCurrencyCode: parsedForeignCurrencyCode,
		Details:             mailRows[7],
		Balance:             float32(parsedBalance),
	})

	return &parsedTransactions
}

func (a *AlfaParser) parseCardNumber(mail string) string {
	re := regexp.MustCompile(cardNumberRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		panic("parse error")
	}

	return match[1]
}

func (a *AlfaParser) parseTime(mail string) time.Time {
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

func (a *AlfaParser) parseTransactionType(value string) string {
	switch value {
	case "Оплата товаров/услуг":
		return entities.ParsedTransactionTypeExpense
	}

	return entities.ParsedTransactionTypeExpense
}

func (a *AlfaParser) parseAmountAndCurrency(value string) (
	amount float64, currency string, foreignAmount float64, foreignCurrency string,
) {
	re := regexp.MustCompile(amountRegexp)
	match := re.FindStringSubmatch(value)

	if match == nil {
		panic("parse error")
	}

	amount = 0.0
	amount, _ = strconv.ParseFloat(match[1], 64)

	foreignAmount = 0.0
	foreignAmount, _ = strconv.ParseFloat(match[4], 64)

	return amount, match[2], foreignAmount, match[5]
}

func (a *AlfaParser) parseBalance(mail string) float64 {
	re := regexp.MustCompile(balanceRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		return 0.0
	}

	balance, _ := strconv.ParseFloat(match[1], 64)

	return balance
}

func (a *AlfaParser) parseDetails(mail string) string {
	re := regexp.MustCompile(detailsRegexp)
	match := re.FindStringSubmatch(mail)

	if match == nil {
		panic("parse error")
	}

	return match[1]
}
