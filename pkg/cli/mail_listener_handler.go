package cli

import (
	"context"
	"time"

	"github.com/getsentry/sentry-go"
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

	err := sentry.Init(sentry.ClientOptions{
		Dsn: settings.SentryDsn,
		Environment: settings.AppEnv,
	})
	if err != nil {
		logger.Fatalf("sentry.Init error: %s", err)
	}
	//defer sentry.Recover()
	defer sentry.Flush(2 * time.Second)

	storage, _ := newDBClient(settings)
	fireflyClient, _ := newFireflyClient(settings)
	telegramBot, _ := newTelegramBot(storage, logger, fireflyClient, settings)
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

	logger.Println("Mail listener started")
	processor.Start(ctx)
}
