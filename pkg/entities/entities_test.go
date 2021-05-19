package entities_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/randaalex/finance_bot/pkg/entities"
)

func TestEntities_Hash(t *testing.T) {
	t.Run("returns hash value", func(t *testing.T) {
		time, _ := time.ParseInLocation("02.01.2006 15:04:05", "15.04.2021 17:36:25", entities.DefaultTZ())

		transaction := entities.ParsedTransaction{
			CardNumber: "1234",
			Time:       time,
			Amount:     100.01,
		}

		assert.Equal(t, "1f1c4da3c53b3d36e669ceb4f6e7a8ab975320add4fcf51d0b9935059aadc471", transaction.Hash())
	})
}
