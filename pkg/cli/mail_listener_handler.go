package cli

import (
	"context"
	"time"

	"github.com/spf13/cobra"

	"github.com/randaalex/finance_bot/pkg/emailprocessor"
	"github.com/randaalex/finance_bot/pkg/emailprocessor/alfaparser"
)

func MailListenerHandler(cmd *cobra.Command, args []string) {
	runMailHandler()
	for {
		time.Sleep(10 * time.Second)
	}
}

func runMailHandler() {
	settings := getSettings()
	logger := newLogger()

	initSentry(logger, settings)

	storage, _ := newDBClient(settings)
	fireflyClient, _ := newFireflyClient(settings)
	telegramBot, _ := newTelegramBot(storage, fireflyClient, settings)
	emailParser := alfaparser.NewParser(logger, getAccounts(), getCategories())

	processor := emailprocessor.NewProcessor(
		storage,
		fireflyClient,
		telegramBot,
		emailParser,
		logger,
		getAccounts(),
		getCategories(),
		settings,
	)

	ctx := context.TODO()

	//err := processor.Connect(ctx)
	//if err != nil {
	//	panic(err) // TODO: fix panic
	//}
	err := processor.Start(ctx)
	if err != nil {
		panic(err) // TODO: fix panic
	}
}
