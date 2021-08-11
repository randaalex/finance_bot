package cli

import (
	"github.com/spf13/cobra"
)

func TelegramBotHandler(cmd *cobra.Command, args []string) {
	settings := getSettings()

	storage, _ := newDBClient(settings)
	fireflyClient, _ := newFireflyClient(settings)

	bot, _ := newTelegramBot(storage, fireflyClient, settings)
	//
	//time, _ := time.ParseInLocation("02.01.2006 15:04:05", "15.04.2021 17:36:25", entities.DefaultTZ())
	//transaction := &entities.FireflyTransaction{
	//	Id:                  123,
	//	AccountId:           1,
	//	AccountName:         "AlfaBank N1 (BYN)",
	//	AccountCurrencyCode: "BYN",
	//	Time:                time,
	//	Type:                "expense",
	//	Amount:              100.01,
	//	ForeignAmount:       0.0,
	//	ForeignCurrencyCode: "USD",
	//	CategoryId:          4,
	//	CategoryName:        "🛒 Food Groceries",
	//	Description:         "BLR/INTERNET/BANK DABRABYT",
	//}
	//bot.RenderTransaction(transaction)

	bot.Start()
}

