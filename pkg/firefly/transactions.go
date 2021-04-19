package firefly

import (
	"context"
	"net/http"
)

type TransactionsServiceInterface interface {
  CreateTransaction(ctx context.Context, request CreateTransactionReq) (*Transaction, error)
}
var _ TransactionsServiceInterface = &TransactionsService{}

type TransactionsService service

type TransactionSubTransaction struct {
	User                 int    `json:"user"`
	TransactionJournalID int    `json:"transaction_journal_id"`
	Type                 string `json:"type"`
	Date                 string `json:"date"`
	Amount               string `json:"amount"`
	Description          string `json:"description"`
	Order                int    `json:"order"`
}
type TransactionAttributes struct {
	CreatedAt  string                      `json:"created_at"`
	Attributes []TransactionSubTransaction `json:"transactions"`
}
type TransactionData struct {
	Type       string                `json:"type"`
	ID         string                `json:"id"`
	Attributes TransactionAttributes `json:"attributes"`
}
type Transaction struct {
	Data TransactionData `json:"data"`
}

type CreateSubTransactionReq struct {
	Type string `json:"type,omitempty"`
	//Date                *Timestamp `json:"date,omitempty"`
	Date                string     `json:"date,omitempty"`
	Amount              float32    `json:"amount,omitempty"`
	Description         string     `json:"description,omitempty"`
	Order               int        `json:"order,omitempty"`
	CurrencyID          int        `json:"currency_id,omitempty"`
	CurrencyCode        string     `json:"currency_code,omitempty"`
	ForeignAmount       float32    `json:"foreign_amount,omitempty"`
	ForeignCurrencyID   int        `json:"foreign_currency_id,omitempty"`
	ForeignCurrencyCode string     `json:"foreign_currency_code,omitempty"`
	BudgetID            int        `json:"budget_id,omitempty"`
	CategoryID          int        `json:"category_id,omitempty"`
	CategoryName        string     `json:"category_name,omitempty"`
	SourceID            int        `json:"source_id,omitempty"`
	SourceName          string     `json:"source_name,omitempty"`
	DestinationID       int        `json:"destination_id,omitempty"`
	DestinationName     string     `json:"destination_name,omitempty"`
	Reconciled          string     `json:"reconciled,omitempty"`
	PiggyBankID         int        `json:"piggy_bank_id,omitempty"`
	PiggyBankName       string     `json:"piggy_bank_name,omitempty"`
	BillId              int        `json:"bill_id,omitempty"`
	BillName            string     `json:"bill_name,omitempty"`
	Tags                []string   `json:"tags,omitempty"`
	Notes               string     `json:"notes,omitempty"`
	InternalReference   string     `json:"internal_reference,omitempty"`
	ExternalID          string     `json:"external_id,omitempty"`
	BunqPaymentID       string     `json:"bunq_payment_id,omitempty"`
	SepaCc              string     `json:"sepa_cc,omitempty"`
	SepaCtOp            string     `json:"sepa_ct_op,omitempty"`
	SepaCtID            string     `json:"sepa_ct_id,omitempty"`
	SepaDb              string     `json:"sepa_db,omitempty"`
	SepaCountry         string     `json:"sepa_country,omitempty"`
	SepaEp              string     `json:"sepa_ep,omitempty"`
	SepaCi              string     `json:"sepa_ci,omitempty"`
	SepaBatchID         string     `json:"sepa_batch_id,omitempty"`
	InterestDate        *Timestamp `json:"interest_date,omitempty"`
	BookDate            *Timestamp `json:"book_date,omitempty"`
	ProcessDate         *Timestamp `json:"process_date,omitempty"`
	DueDate             *Timestamp `json:"due_date,omitempty"`
	PaymentDate         *Timestamp `json:"payment_date,omitempty"`
	InvoiceDate         *Timestamp `json:"invoice_date,omitempty"`
}

type CreateTransactionReq struct {
	ErrorIfDuplicateHash bool   `json:"error_if_duplicate_hash,omitempty"`
	ApplyRules           bool   `json:"apply_rules,omitempty"`
	GroupTitle           string `json:"group_title,omitempty"`

	Transactions []CreateSubTransactionReq `json:"transactions"`
}

func (s *TransactionsService) CreateTransaction(
	ctx context.Context,
	request CreateTransactionReq,
) (*Transaction, error) {
	const url = "transactions"

	req, err := s.client.NewRequest(http.MethodPost, url, request)
	if err != nil {
		return nil, err
	}

	var transaction Transaction

	_, err = s.client.Do(ctx, req, &transaction)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("===> Create transaction in firefly: %v", &request)

	return &transaction, nil
}

//func (s *TransactionsService) UpdateTransaction(
//	ctx context.Context,
//	request CreateTransactionReq,
//) (*Transaction, *Response, error) {
//
//}
