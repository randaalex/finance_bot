package fileprocessor

import (
	"context"
	"fmt"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
)

type Storage interface {
	CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error)
	GetMappingByTransactionDetails(ctx context.Context, transactionDetails string) (db.Mapping, error)
}

type Processor struct {
	Storage          Storage
	FireflyClient    *firefly.Client
	CategorySelector CategorySelectorInterface
}

func NewProcessor(storage Storage, fireflyClient *firefly.Client, categorySelector CategorySelectorInterface) *Processor {
	return &Processor{
		Storage:          storage,
		FireflyClient:    fireflyClient,
		CategorySelector: categorySelector,
	}
}

func (p *Processor) Process(transactions *[]entities.ParsedTransaction) {
	for _, transaction := range *transactions {
		p.processTransaction(&transaction)
	}
}

func (p *Processor) processTransaction(transaction *entities.ParsedTransaction) {
	ctx := context.TODO()
	preselectedCategoryID := 0

	existedTransaction, e := p.Storage.GetMappingByTransactionDetails(ctx, transaction.Details)
	if e != nil {
		fmt.Printf("transaction not found: %v\n", transaction.Hash())
	} else {
		fmt.Printf("transaction found: %v\n", existedTransaction)
		preselectedCategoryID = int(existedTransaction.CategoryID)
	}

	fmt.Printf("preselected category: %v\n", preselectedCategoryID)

	selectedCategoryID := p.CategorySelector.Select(transaction, preselectedCategoryID)

	fmt.Printf("selected category: %v\n", selectedCategoryID)

	_, err := p.Storage.CreateMapping(ctx, db.CreateMappingParams{
		TransactionDetails: transaction.Details,
		CategoryID:         int32(selectedCategoryID),
	})
	if err != nil {
		panic(err)
	}

	p.createTransactionInFirefly(
		ctx,
		transaction,
		selectedCategoryID,
	)
}

func (p *Processor) createTransactionInFirefly(
	ctx context.Context,
	transaction *entities.ParsedTransaction,
	selectedCategoryID int,
) *firefly.Transaction {
	req := firefly.CreateTransactionReq{
		Transactions: []firefly.CreateSubTransactionReq{
			{
				Type:          "deposit",
				Date:          transaction.Time.Format("2006-01-02"),
				Amount:        transaction.Amount,
				Description:   transaction.Details,
				SourceID:      4,
				DestinationID: selectedCategoryID,
				Notes:         transaction.Details,
				ExternalID:    transaction.Hash(),
			},
		},
	}

	createdTransaction, err := p.FireflyClient.Transactions.CreateTransaction(ctx, req)
	fmt.Println("###########")
	fmt.Println(createdTransaction, err)
	fmt.Println("###########")
	return createdTransaction
}
