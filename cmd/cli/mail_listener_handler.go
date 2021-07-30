package cli

import (
	"context"
	"github.com/randaalex/finance_bot/pkg/parsers/alfamail"
	"time"

	"github.com/spf13/cobra"

	"github.com/randaalex/finance_bot/pkg/processors/emailprocessor"
)

func MailListenerHandler(cmd *cobra.Command, args []string) {
	runMailHandler()
	for {
		time.Sleep(10 * time.Second)
	}
}

func runMailHandler() {
	settings := getSettings()

	storage, _ := newDBClient(settings)
	fireflyClient, _ := newFireflyClient(settings)
	telegramBot, _ := newTelegramBot(storage, fireflyClient, settings)
	emailParser := alfamail.NewAlfaParser(nil, getAccounts(), getCategories())

	processor := emailprocessor.NewProcessor(
		storage,
		fireflyClient,
		telegramBot,
		emailParser,
		getAccounts(),
		getCategories(),
		&emailprocessor.Settings{
			Address:  settings.MailClientAddress,
			Username: settings.MailClientUsername,
			Password: settings.MailClientPassword,
		},
	)

	ctx := context.TODO()

	processor.DownloadEmails(ctx)
	processor.Process(ctx)
}
