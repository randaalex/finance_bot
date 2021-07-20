package bot

import (
	"context"
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
	FireflyClient *firefly.Client
	Bot           *telebot.Bot
}

func NewBot(storage Storage, fireflyClient *firefly.Client) *Bot {
	settings := telebot.Settings{
		Token:  "token",
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	bot, err := telebot.NewBot(settings)

	if err != nil {
		panic(err)
	}

	return &Bot{
		Storage:       storage,
		FireflyClient: fireflyClient,
		Bot:           bot,
	}
}

func (b *Bot) Start(storage Storage, fireflyClient *firefly.Client) {
	b.Bot.Handle("/hello", func(m *telebot.Message) {
		b.Bot.Send(m.Sender, "Hello World!")
	})

	b.Bot.Start()
}
