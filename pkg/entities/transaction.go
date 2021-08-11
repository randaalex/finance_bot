package entities

import (
	"crypto/md5"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"strconv"
	"time"
)

const (
	TransactionSplitTypeWithdrawal = "withdrawal"
	TransactionSplitTypeDeposit    = "deposit"
	TransactionSplitTypeTransfer   = "transfer"

	FireflyTransactionAmountFormat   = "%.2f"
	FireflyTransactionDateFormat     = "2006-01-02"
	FireflyTransactionDateTimeFormat = "2006.01.02 15:04"
)

type Transaction struct {
	Id                  int
	IssuedAt            time.Time // Internal Reference
	Type                string
	Description         string // custom or ExternalId
	Amount              float64
	CurrencyCode        string
	ForeignAmount       float64
	ForeignCurrencyCode string
	CategoryId          int
	CategoryName        string
	SourceId            int
	SourceName          string
	DestinationId       int
	DestinationName     string
	Date                string
	Recipient           string // firefly: ExternalId
	Notes               string // full email content
}

func (t *Transaction) GetHash() string {
	uniqValue := fmt.Sprintf("%v-%v-%v-%v", t.IssuedAt, t.SourceId, t.DestinationId, t.Amount)

	hash := md5.Sum([]byte(uniqValue))

	return fmt.Sprintf("%x", hash)
}

func ConvertTransactionToFireflyTransaction(transaction *Transaction) *firefly.Transaction {
	errorIfDuplicateHash := true

	hash := transaction.GetHash()
	internalReference := transaction.IssuedAt.Format(FireflyTransactionDateTimeFormat)

	fireflySplit := firefly.TransactionSplit{
		Date:                transaction.IssuedAt.Format(FireflyTransactionDateFormat),
		Type:                transaction.Type,
		Description:         transaction.Description,
		Amount:              fmt.Sprintf(FireflyTransactionAmountFormat, transaction.Amount),
		CategoryId:          *NewNullableInt32(&transaction.CategoryId),
		SourceId:            *NewNullableInt32(&transaction.SourceId),
		SourceName:          *NewNullableString(&transaction.SourceName),
		DestinationId:       *NewNullableInt32(&transaction.DestinationId),
		DestinationName:     *NewNullableString(&transaction.DestinationName),
		ExternalId:          *NewNullableString(&transaction.Recipient),
		ImportHashV2:        *NewNullableString(&hash),
		InternalReference:   *NewNullableString(&internalReference),
		Notes:               *NewNullableString(&transaction.Notes),
	}

	if transaction.ForeignAmount > 0 {
		fireflySplit.SetForeignAmount(fmt.Sprintf(FireflyTransactionAmountFormat, transaction.ForeignAmount))
		fireflySplit.SetForeignCurrencyCode(transaction.ForeignCurrencyCode)
	}

	fireflyTransaction := firefly.Transaction{
		ErrorIfDuplicateHash: &errorIfDuplicateHash,
		Transactions: []firefly.TransactionSplit{fireflySplit},
	}

	return &fireflyTransaction
}

func ConvertFireflyTransactionToTransaction(fireflyTransactionSingle *firefly.TransactionSingle) *Transaction {
	fireflyTransactionRead := fireflyTransactionSingle.GetData()
	fireflyTransaction := fireflyTransactionRead.GetAttributes()
	fireflyTransactionSplits := fireflyTransaction.GetTransactions()

	transactionId, _ := strconv.Atoi(fireflyTransactionRead.GetId())

	if len(fireflyTransactionSplits) == 0 {
		panic("empty firefly transactions split") // TODO: fix panic
	}

	// we use only one TransactionSplit in firefly
	fireflySplit := fireflyTransactionSplits[0]

	amount, _ := strconv.ParseFloat(fireflySplit.GetAmount(), 32)
	foreignAmount, _ := strconv.ParseFloat(fireflySplit.GetForeignAmount(), 32)
	issuedAt, err := time.ParseInLocation(FireflyTransactionDateTimeFormat, fireflySplit.GetInternalReference(), DefaultTZ())
	if err != nil {
		issuedAt = time.Time{}
	}

	transaction := Transaction{
		Id:                  transactionId,
		IssuedAt:            issuedAt,
		Type:                fireflySplit.GetType(),
		Description:         fireflySplit.GetDescription(),
		Amount:              amount,
		CurrencyCode:        fireflySplit.GetCurrencyCode(),
		ForeignAmount:       foreignAmount,
		ForeignCurrencyCode: fireflySplit.GetForeignCurrencyCode(),
		CategoryId:          int(fireflySplit.GetCategoryId()),
		CategoryName:        fireflySplit.GetCategoryName(),
		SourceId:            int(fireflySplit.GetSourceId()),
		SourceName:          fireflySplit.GetSourceName(),
		DestinationId:       int(fireflySplit.GetDestinationId()),
		DestinationName:     fireflySplit.GetDestinationName(),
		Recipient:           fireflySplit.GetExternalId(),
		Notes:               fireflySplit.GetNotes(),
	}

	return &transaction
}
