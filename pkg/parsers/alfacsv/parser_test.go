package alfacsv_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"

	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/parsers/alfacsv"
)

func TestAlfacsv_Parse(t *testing.T) {
	parser := alfacsv.NewAlfaParser(nil)

	var emptyResult []entities.ParsedTransaction

	var encodeInput = func(input string) bytes.Buffer {
		var (
			reader bytes.Buffer
			buffer bytes.Buffer
		)

		encoder := transform.NewWriter(&buffer, charmap.Windows1251.NewEncoder())
		encoder.Write([]byte(input))
		encoder.Close()

		reader.WriteString(buffer.String())

		return reader
	}

	t.Run("when file is empty", func(t *testing.T) {
		reader := encodeInput("")

		t.Run("returns empty result", func(t *testing.T) {
			result := parser.Parse(&reader)

			assert.Equal(t, &emptyResult, result)
		})
	})

	t.Run("when file has invalid line", func(t *testing.T) {
		reader := encodeInput("invalid")

		t.Run("returns empty result", func(t *testing.T) {
			result := parser.Parse(&reader)

			assert.Equal(t, &emptyResult, result)
		})
	})

	t.Run("when file has valid transaction line", func(t *testing.T) {
		reader := encodeInput("99223254647858;15.04.2021 17:36:25;Безналичная операция;Завершено успешно;209.00;BYN;MINSK;BLR;14913.SHOP.ONLINER.BY")

		t.Run("returns valid parsed transaction", func(t *testing.T) {
			result := parser.Parse(&reader)

			expectedTime, _ := time.ParseInLocation("02.01.2006 15:04:05", "15.04.2021 17:36:25", entities.DefaultTZ())

			expectedResult := []entities.ParsedTransaction{
				{
					CardNumber:          "7858",
					Time:                expectedTime,
					Type:                entities.ParsedTransactionTypeExpense,
					Amount:              209,
					CurrencyCode:        "BYN",
					ForeignAmount:       0,
					ForeignCurrencyCode: "",
					Details:             "MINSK/BLR/14913.SHOP.ONLINER.BY",
					Balance:             0,
				},
			}

			assert.Equal(t, &expectedResult, result)
		})
	})
}
