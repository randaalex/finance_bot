package entities

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type ParsedTransaction struct {
	CardNumber          string
	Time                time.Time
	Type                string
	Amount              float32
	CurrencyCode        string
	ForeignAmount       float32
	ForeignCurrencyCode string
	Details             string
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
