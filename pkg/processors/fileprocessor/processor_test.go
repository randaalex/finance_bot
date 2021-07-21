package fileprocessor_test

import (
	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	fireflyMocks "github.com/randaalex/finance_bot/pkg/mocks/firefly"
	fileprocessorMocks "github.com/randaalex/finance_bot/pkg/mocks/processors/fileprocessor"
	"github.com/randaalex/finance_bot/pkg/processors/fileprocessor"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
	"time"
)

func TestFileprocessor_Process(t *testing.T) {
	storageMock := fileprocessorMocks.Storage{}
	transactionsFireflyClientMock := fireflyMocks.TransactionsServiceInterface{}

	fireflyClient := firefly.NewClient(&http.Client{}, "")
	fireflyClient.Transactions = &transactionsFireflyClientMock

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

			transactionsFireflyClientMock.
				On("CreateTransaction",
					mock.Anything,
					firefly.CreateTransactionReq{
						Transactions: []firefly.CreateSubTransactionReq{
							{
								Type:          "deposit",
								Date:          "2021-04-15",
								Amount:        100,
								Description:   "details",
								SourceID:      4,
								DestinationID: 0,
								Notes:         "details",
								ExternalID:    "38b86679585cca16fb0f827559ee5ff3dff4f97666593eafbf47d478be37a76d",
							},
						},
					},
				).
				Return(&firefly.Transaction{}, nil)

			processor.Process(&transactions)
		})
	})
}