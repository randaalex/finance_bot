package fileprocessor_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	fireflyMocks "github.com/randaalex/finance_bot/pkg/mocks/firefly"
	fileprocessorMocks "github.com/randaalex/finance_bot/pkg/mocks/processors/fileprocessor"
	"github.com/randaalex/finance_bot/pkg/processors/fileprocessor"
)

func TestFileprocessor_Process(t *testing.T) {
	storageMock := fileprocessorMocks.Storage{}
	fireflyTransactionsApiMock := fireflyMocks.TransactionsApi{}

	fireflyClient := firefly.NewAPIClient(&firefly.Configuration{})
	fireflyClient.TransactionsApi = &fireflyTransactionsApiMock

	categorySelectorMock := fileprocessorMocks.CategorySelectorInterface{}

	processor := fileprocessor.NewProcessor(&storageMock, fireflyClient, &categorySelectorMock)

	t.Run("successfully process", func(t *testing.T) {
		time, _ := time.ParseInLocation("02.01.2006 15:04:05", "15.04.2021 17:36:25", entities.DefaultTZ())

		transaction := &entities.ParsedTransaction{
			CardNumber:          "1234",
			Time:                time,
			Type:                "type",
			Amount:              100.00,
			CurrencyCode:        "BYN",
			ForeignAmount:       200.00,
			ForeignCurrencyCode: "USD",
			Details:             "details",
		}

		transactions := []entities.ParsedTransaction{*transaction}

		t.Run("transaction not exists in DB", func(t *testing.T) {
			storageMock.
				On("GetMappingByTransactionDetails", mock.Anything, "details").
				Return(db.Mapping{}, nil)

			storageMock.
				On("CreateMapping",
					mock.Anything,
					db.CreateMappingParams{
						TransactionDetails: "details",
						CategoryID:         0,
					}).
				Return(db.Mapping{}, nil)

			categorySelectorMock.
				On("Select",
					transaction,
					0,
				).
				Return(0)

			fireflyTransactionsApiMock.
				On("StoreTransaction", mock.Anything).
				Return(firefly.ApiStoreTransactionRequest{
					ApiService: &fireflyTransactionsApiMock,
				}, nil)

			// TODO: check actual request body instead of mock.Anything
			fireflyTransactionsApiMock.
				On("StoreTransactionExecute", mock.Anything).
				Return(firefly.TransactionSingle{}, nil, nil)

			processor.Process(&transactions)
		})
	})
}
