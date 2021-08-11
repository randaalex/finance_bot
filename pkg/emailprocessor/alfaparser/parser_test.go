package alfaparser_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"

	"github.com/randaalex/finance_bot/pkg/emailprocessor/alfaparser"
	"github.com/randaalex/finance_bot/pkg/entities"
)

func TestAlfamail_Parse(t *testing.T) {
	accounts := &[]entities.Account{
		{
			Id:            1,
			Name:          "AlfaBank N1 (BYN)",
			CurrencyCode:  "BYN",
			AccountNumber: "1111",
		},
	}

	categories := &[]entities.Category{
		{Id: 1, Name: "cat1"},
		{Id: 2, Name: "cat2"},
	}

	parser := alfaparser.NewParser(nil, accounts, categories)

	t.Run("when email is empty", func(t *testing.T) {
		data := ""

		t.Run("returns invalid email error", func(t *testing.T) {
			result, err := parser.Parse(data)

			assert.Error(t, err, "invalid mail error")
			assert.Nil(t, result)
		})
	})

	t.Run("when email has invalid content", func(t *testing.T) {
		data := "invalid"

		t.Run("returns match account error", func(t *testing.T) {
			result, err := parser.Parse(data)

			assert.Error(t, err, "match account error")
			assert.Nil(t, result)
		})
	})

	t.Run("when email has valid transaction", func(t *testing.T) {
		data := `Карта 1.1111
Со счёта: BY06ALFA30143400080000000000
Оплата товаров/услуг
Успешно
Сумма:10.20 USD (20.41 BYN)
Остаток:500.01 BYN
На время:17:36:25
MINSK/BLR/14913.SHOP.ONLINER.BY
15.04.2021 17:36:25`

		t.Run("returns valid parsed transaction", func(t *testing.T) {
			result, _ := parser.Parse(data)

			expectedTime, _ := time.ParseInLocation("02.01.2006 15:04:05", "15.04.2021 17:36:25", entities.DefaultTZ())

			expectedResult := entities.Transaction{
				Id:                  0,
				IssuedAt:            expectedTime,
				Type:                entities.TransactionSplitTypeWithdrawal,
				Description:         "MINSK/BLR/14913.SHOP.ONLINER.BY",
				Amount:              20.41,
				CurrencyCode:        "BYN",
				ForeignAmount:       10.2,
				ForeignCurrencyCode: "USD",
				CategoryId:          0,
				CategoryName:        "",
				SourceId:            1,
				SourceName:          "",
				DestinationId:       0,
				DestinationName:     "",
				Date:                "",
				Recipient:           "MINSK/BLR/14913.SHOP.ONLINER.BY",
				Notes:               data,
			}

			assert.Equal(t, &expectedResult, result)
		})
	})

	t.Run("when email has valid deposit transaction", func(t *testing.T) {
		data := `Карта 4.1111
На счёт: BY06ALFA30143400080000000000
Поступление
Успешно
Сумма:200.00 BYN
Остаток:203.78 BYN
На время:17:36:25
15.04.2021`

		t.Run("returns valid parsed transaction", func(t *testing.T) {
			result, err := parser.Parse(data)

			expectedTime, _ := time.ParseInLocation("02.01.2006 15:04:05", "15.04.2021 17:36:25", entities.DefaultTZ())

			expectedResult := entities.Transaction{
				Id:                  0,
				IssuedAt:            expectedTime,
				Type:                entities.TransactionSplitTypeDeposit,
				Description:         "income",
				Amount:              200.0,
				CurrencyCode:        "BYN",
				ForeignAmount:       0,
				ForeignCurrencyCode: "",
				CategoryId:          0,
				CategoryName:        "",
				SourceId:            0,
				SourceName:          "",
				DestinationId:       1,
				DestinationName:     "",
				Date:                "",
				Recipient:           "",
				Notes:               data,
			}

			assert.NoError(t, err)
			assert.Equal(t, &expectedResult, result)
		})
	})
}
