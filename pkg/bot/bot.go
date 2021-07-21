package bot

import (
	"context"
	"github.com/randaalex/finance_bot/pkg/entities"
	"time"

	"gopkg.in/tucnak/telebot.v2"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/firefly"
)

type Storage interface {
	CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error)
	GetMappingByTransactionDetails(ctx context.Context, transactionDetails string) (db.Mapping, error)
}

type Bot struct {
	Storage       Storage
	FireflyClient *firefly.APIClient
	Telebot       *telebot.Bot
	User          *telebot.User
}

func NewBot(storage Storage, fireflyClient *firefly.APIClient, settings *entities.Settings) *Bot {
	botSettings := telebot.Settings{
		Token:  settings.TelegramBotToken,
		Poller: &telebot.LongPoller{
			Timeout: settings.TelegramPollingInterval * time.Second,
		},
	}

	bot, err := telebot.NewBot(botSettings)

	if err != nil {
		panic(err)
	}

	return &Bot{
		Storage:       storage,
		FireflyClient: fireflyClient,
		Telebot:       bot,
		User:          &telebot.User{ID: settings.TelegramUserId},
	}
}

func (b *Bot) Start() {
	b.Telebot.Handle(&telebot.InlineButton{Unique: btnUpdateTransactionCategory}, b.updateTransactionCategoryHandler)
	b.Telebot.Handle(&telebot.InlineButton{Unique: btnDeleteTransaction}, b.deleteTransactionHandler)

	b.Telebot.Start()
}
