package entities

import "time"

type Settings struct {
	AppEnv string `mapstructure:"APP_ENV"`

	DbConnection string `mapstructure:"DB_CONNECTION"`

	FireflyHost   string `mapstructure:"FIREFLY_HOST"`
	FireflyScheme string `mapstructure:"FIREFLY_SCHEME"`
	FireflyToken  string `mapstructure:"FIREFLY_TOKEN"`

	MailClientAddress  string `mapstructure:"MAIL_CLIENT_ADDRESS"`
	MailClientUsername string `mapstructure:"MAIL_CLIENT_USERNAME"`
	MailClientPassword string `mapstructure:"MAIL_CLIENT_PASSWORD"`

	TelegramBotToken        string        `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramPollingInterval time.Duration `mapstructure:"TELEGRAM_POLLING_INTERVAL"`
	TelegramUserId          int           `mapstructure:"TELEGRAM_USER_ID"`

	SentryDsn string `mapstructure:"SENTRY_DSN"`
}
