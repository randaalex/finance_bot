package fileprocessor

import (
	"context"
	"fmt"

	"github.com/manifoldco/promptui"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
)

type Storage interface {
	CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error)
	GetMappingByTransactionDetails(ctx context.Context, transactionDetails string) (db.Mapping, error)
}

type Processor struct {
	Storage       Storage
	FireflyClient *firefly.Client
}

func NewProcessor(storage Storage, fireflyClient *firefly.Client) *Processor {
	return &Processor{
		Storage:       storage,
		FireflyClient: fireflyClient,
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

	selectedCategoryID := p.askCategory(transaction, preselectedCategoryID)

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

func (p *Processor) askCategory(transaction *entities.ParsedTransaction, preselectedCategoryID int) int {
	type category struct {
		ID   int
		Name string
	}

	categories := []category{
		{ID: 3, Name: "Work"},
		{ID: 4, Name: "Food Grocering"},
		{ID: 5, Name: "Food Restaurants"},
		{ID: 1, Name: "House Bills"},
		{ID: 6, Name: "House Purchases"},
		{ID: 7, Name: "House Pets"},
		{ID: 8, Name: "House Credit"},
		{ID: 9, Name: "Leisure General"},
		{ID: 2, Name: "Leisure Tourism"},
		{ID: 10, Name: "Leisure Hobby"},
		{ID: 11, Name: "Leisure Presents"},
		{ID: 12, Name: "Leisure Entertainment"},
		{ID: 13, Name: "Personal"},
		// {ID: 13, Name: "Health"},
		{ID: 14, Name: "Car Fuel"},
		{ID: 15, Name: "Car Maintenance"},
		{ID: 16, Name: "Clothing"},
		{ID: 17, Name: "Other"},
		{ID: 18, Name: "Corrections"},
	}

	preselectedCategoryNumber := 0

	for index, category := range categories {
		if category.ID == preselectedCategoryID {
			preselectedCategoryNumber = index
			break
		}
	}

	prompt := promptui.Select{
		Label:             fmt.Sprintf("Select category for '%v'", transaction),
		Items:             categories,
		StartInSearchMode: false,
		//Searcher: func(input string, index int) bool {
		//	return strings.Contains(categories[index], input)
		//},
		Templates: &promptui.SelectTemplates{
			Active:   "=> {{ .ID | green }} {{ .Name | green }}",
			Inactive: "   {{ .ID }} {{ .Name }}",
			Selected: "{{ .ID }}",
		},
		CursorPos:    preselectedCategoryNumber,
		HideSelected: true,
		Size:         99,
	}
	selectedCategoryNumber, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return 0
	}

	selectedCategory := categories[selectedCategoryNumber]

	return selectedCategory.ID
}
