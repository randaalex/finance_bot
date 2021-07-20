package alfamail_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/parsers/alfamail"
)

func TestAlfamail_Parse(t *testing.T) {
	parser := alfamail.NewAlfaParser(nil)

	var emptyResult []entities.ParsedTransaction

	t.Run("when email is empty", func(t *testing.T) {
		data := ""

		t.Run("returns empty result", func(t *testing.T) {
			result := parser.Parse(data)

			assert.Equal(t, &emptyResult, result)
		})
	})

	t.Run("when email has invalid content", func(t *testing.T) {
		data := "invalid"

		t.Run("returns empty result", func(t *testing.T) {
			result := parser.Parse(data)

			assert.Equal(t, &emptyResult, result)
		})
	})

	t.Run("when email has valid transaction", func(t *testing.T) {
		data := `Карта 4.7858
Со счёта: BY06ALFA30143400080000000000
Оплата товаров/услуг
Успешно
Сумма:10.20 USD (20.41 BYN)
Остаток:500.01 BYN
На время:17:36:25
MINSK/BLR/14913.SHOP.ONLINER.BY
15.04.2021 17:36:25`

		t.Run("returns valid parsed transaction", func(t *testing.T) {
			result := parser.Parse(data)

			expectedTime, _ := time.ParseInLocation("02.01.2006 15:04:05", "15.04.2021 17:36:25", entities.DefaultTZ())

			expectedResult := []entities.ParsedTransaction{
				{
					CardNumber:          "7858",
					Time:                expectedTime,
					Type:                entities.ParsedTransactionTypeExpense,
					Amount:              10.20,
					CurrencyCode:        "USD",
					ForeignAmount:       20.41,
					ForeignCurrencyCode: "BYN",
					Details:             "MINSK/BLR/14913.SHOP.ONLINER.BY",
					Balance:             500.01,
				},
			}

			assert.Equal(t, &expectedResult, result)
		})
	})
}
