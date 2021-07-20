package entities

import (
	"crypto/sha256"
	"fmt"
	"time"
)

func DefaultTZ() *time.Location {
	tz, _ := time.LoadLocation("Europe/Minsk")

	return tz
}

type ParsedTransaction struct {
	CardNumber          string
	Time                time.Time
	Type                string
	Amount              float32
	CurrencyCode        string
	ForeignAmount       float32
	ForeignCurrencyCode string
	Details             string
	Balance             float32
}

func (p *ParsedTransaction) Hash() string {
	uniqValue := fmt.Sprintf("%v-%v-%v", p.CardNumber, p.Time, p.Amount)
	shaBytes := sha256.Sum256([]byte(uniqValue))

	return fmt.Sprintf("%x", shaBytes)
}

const (
	ParsedTransactionTypeExpense = "expense"
	//ParsedTransactionIncome   = "income"
	//ParsedTransactionTransfer = "transfer"
)

type FireflyTransaction struct {
	ID int
}
