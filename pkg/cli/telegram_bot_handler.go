package cli

import (
	"github.com/getsentry/sentry-go"
	"github.com/spf13/cobra"
	"time"
)

func TelegramBotHandler(cmd *cobra.Command, args []string) {
	settings := getSettings()
	logger := newLogger()

	err := sentry.Init(sentry.ClientOptions{
		Dsn: settings.SentryDsn,
		Environment: settings.AppEnv,
	})
	if err != nil {
		logger.Fatalf("sentry.Init error: %s", err)
	}
	defer sentry.Recover()
	defer sentry.Flush(2 * time.Second)

	storage, _ := newDBClient(settings)
	fireflyClient, _ := newFireflyClient(settings)

	bot, _ := newTelegramBot(storage, logger, fireflyClient, settings)

	logger.Println("Bot started")

	bot.Start()
}

