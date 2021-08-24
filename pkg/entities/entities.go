package entities

import (
	"crypto/sha256"
	"fmt"
	"time"
)

type Account struct {
	Id            int
	Name          string
	CurrencyCode  string
	AccountNumber string
}

type Category struct {
	Type string
	Id   int
	Name string
}

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

// income or expense
type FireflyTransaction struct {
	Id                  int
	AccountId           int
	AccountName         string
	AccountCurrencyCode string
	Time                time.Time
	Type                string  // expense or income
	Amount              float32 // in account currency
	ForeignAmount       float32
	ForeignCurrencyCode string
	CategoryId          int // firefly category id
	CategoryName        string
	Hash                string // store in Internal reference
	Description         string
	RawContent          string // store in notes

	Tags []string // ???
}

func (f *FireflyTransaction) GetCategoryName() string {
	if f.CategoryName == "" {
		return "❔ Without category"
	}

	return f.CategoryName
}
