package alfacsv

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	separator  = ";"
	timeFormat = "02.01.2006 15:04:05"
	rowFormat  = `^\d{14}`
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

func (a *AlfaParser) Parse(reader io.Reader) *[]entities.ParsedTransaction {
	fileReader := transform.NewReader(reader, encoding.NewDecoder())

	var parsedTransactions []entities.ParsedTransaction

	scanner := bufio.NewScanner(fileReader)
	for scanner.Scan() {
		row := scanner.Text()
		fmt.Printf("parse row: %v\n", row)

		parsedRow, parseRowErr := a.parseRow(row)
		if parseRowErr != nil {
			fmt.Printf("parse row error: %v | %v\n", parseRowErr, row)
			continue
		}

		parsedTransactions = append(parsedTransactions, *parsedRow)
	}

	return &parsedTransactions
}

func (a *AlfaParser) parseRow(row string) (*entities.ParsedTransaction, error) {
	isTransactionRow, _ := regexp.Match(rowFormat, []byte(row))

	if !isTransactionRow {
		return nil, errors.New("invalid row format")
	}

	parsedRow := strings.Split(row, separator)

	transactionID := parsedRow[0]

	cardNumber := transactionID[len(transactionID)-4:]
	parsedTime, _ := time.ParseInLocation(timeFormat, parsedRow[1], entities.DefaultTZ())
	parsedAmount, _ := strconv.ParseFloat(parsedRow[4], 64)

	return &entities.ParsedTransaction{
		CardNumber:          cardNumber,
		Time:                parsedTime,
		Type:                a.parseTransactionType(parsedRow[2]),
		Amount:              float32(parsedAmount),
		CurrencyCode:        parsedRow[5],
		ForeignAmount:       0,
		ForeignCurrencyCode: "",
		Details:             fmt.Sprintf("%s/%s/%s", parsedRow[6], parsedRow[7], parsedRow[8]),
		Balance:             0,
	}, nil
}

func (a *AlfaParser) parseTransactionType(value string) string {
	switch value {
	case "Безналичная операция", "Списание", "Отправление средств":
		return entities.ParsedTransactionTypeExpense
	}

	return entities.ParsedTransactionTypeExpense
}
