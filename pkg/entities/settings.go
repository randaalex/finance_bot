package entities

import "time"

type Settings struct {
	DbConnection string `mapstructure:"DB_CONNECTION"`

	FireflyBaseUrl string `mapstructure:"FIREFLY_BASE_URL"`
	FireflyToken   string `mapstructure:"FIREFLY_TOKEN"`

	MailClientAddress  string `mapstructure:"MAIL_CLIENT_ADDRESS"`
	MailClientUsername string `mapstructure:"MAIL_CLIENT_USERNAME"`
	MailClientPassword string `mapstructure:"MAIL_CLIENT_PASSWORD"`

	TelegramBotToken        string        `mapstructure:"TELEGRAM_BOT_TOKEN"`
	TelegramPollingInterval time.Duration `mapstructure:"TELEGRAM_POLLING_INTERVAL"`
	TelegramUserId          int           `mapstructure:"TELEGRAM_USER_ID"`
}