package cli

import (
	"time"

	"github.com/spf13/cobra"

	"github.com/randaalex/finance_bot/pkg/processors/emailprocessor"
)

func MailListenerHandler(cmd *cobra.Command, args []string) {
	//c := cron.New()
	//_ = c.AddFunc("*/30 * * * * *", runMailHandler)
	//c.Start()
	//
	//fmt.Println("Cron scheduler started:")
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

	processor := emailprocessor.NewProcessor(
		storage,
		fireflyClient,
		telegramBot,
		getAccounts(),
		getCategories(),
		&emailprocessor.Settings{
			Address:  settings.MailClientAddress,
			Username: settings.MailClientUsername,
			Password: settings.MailClientPassword,
		},
	)

	processor.DownloadEmails()
	processor.Process()
}
