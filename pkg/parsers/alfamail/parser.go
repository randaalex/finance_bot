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
	timeFormat      = "02.01.2006 15:04:05"
	validMailFormat = `^Карта`
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

	if isValid, _ := regexp.Match(validMailFormat, []byte(mail)); !isValid {
		fmt.Printf("parse error: %v | %v\n", errors.New("invalid mail format"), mail)
		return &parsedTransactions
	}

	mailRows := strings.Split(mail, "\n")

	rawCardNumber := mailRows[0]
	parsedCardNumber := rawCardNumber[len(rawCardNumber)-4:]

	parsedTime, _ := time.ParseInLocation(timeFormat, mailRows[len(mailRows)-1], entities.DefaultTZ())

	parsedType := a.parseTransactionType(mailRows[2])
	parsedAmount, parsedCurrencyCode, parsedForeignAmount, parsedForeignCurrencyCode := a.parseAmountAndCurrency(mailRows[4])

	parsedBalance := a.parseBalance(mailRows[5])

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
	re := regexp.MustCompile(`^Сумма:(\d+\.?\d{0,2})\s(\w{3})(\s\((\d+\.?\d{0,2})\s(\w{3})\))?$`)
	match := re.FindStringSubmatch(value)

	amount = 0.0
	amount, _ = strconv.ParseFloat(match[1], 64)

	foreignAmount = 0.0
	foreignAmount, _ = strconv.ParseFloat(match[4], 64)

	return amount, match[2], foreignAmount, match[5]
}

func (a *AlfaParser) parseBalance(value string) float64 {
	re := regexp.MustCompile(`^Остаток:(\d+\.?\d{0,2})\s(\w{3})$`)
	match := re.FindStringSubmatch(value)

	balance := 0.0
	balance, _ = strconv.ParseFloat(match[1], 64)

	return balance
}
