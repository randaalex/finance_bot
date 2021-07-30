package cli

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"github.com/randaalex/finance_bot/pkg/telegram"
)

type Storage interface {
	CreateTransactionsLog(ctx context.Context, arg db.CreateTransactionsLogParams) (db.TransactionsLog, error)
	GetTransactionsLogByDescription(ctx context.Context, description string) (db.TransactionsLog, error)
}

func newDBClient(settings *entities.Settings) (Storage, error) {
	dbConnection, err := sql.Open("postgres", settings.DbConnection)
	CheckError(err)

	err = dbConnection.Ping()
	CheckError(err)

	dbStorage := db.New(dbConnection)

	return dbStorage, nil
}

func newFireflyClient(settings *entities.Settings) (*firefly.APIClient, error) {
	ctx := context.TODO()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: settings.FireflyToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	return firefly.NewAPIClient(&firefly.Configuration{
		Host:       "192.168.0.2:8081",
		Scheme:     "http",
		HTTPClient: tc,
		OperationServers: map[string]firefly.ServerConfigurations{
			"TransactionsApiService.StoreTransaction":  {{URL: ""}},
			"TransactionsApiService.GetTransaction":    {{URL: ""}},
			"TransactionsApiService.UpdateTransaction": {{URL: ""}},
		},
	}), nil
}

func newTelegramBot(storage Storage, fireflyClient *firefly.APIClient, settings *entities.Settings) (*telegram.Bot, error) {
	return telegram.NewBot(storage, fireflyClient, settings, getAccounts(), getCategories()), nil
}

func getSettings() *entities.Settings {
	var settings entities.Settings

	_ = viper.Unmarshal(&settings)

	return &settings
}

func getAccounts() *[]entities.Account {
	// TODO: load from API
	return &[]entities.Account{
		{
			Id: 7,
			Name: "AlfaBank N1 (BYN)",
			CurrencyCode: "BYN",
			AccountNumber: "7858",
		},
		{
			Id: 8,
			Name: "AlfaBank N1 (USD)",
			CurrencyCode: "USD",
			AccountNumber: "6185",
		},
	}
}

func getCategories() *[]entities.Category {
	// TODO: load from API
	return &[]entities.Category{
		{ Id: 14, Name: "⛽️ Car Fuel" },
		{ Id: 15, Name: "🛠 Car Maintenance" },
		{ Id: 4, Name:  "🛒 Food Groceries" },
		{ Id: 5, Name:  "🍔 Food Restaurants" },
		{ Id: 1, Name:  "🏠 House Bills" },
		{ Id: 6, Name:  "🏠 House Purchases" },
		{ Id: 7, Name:  "🐈 House Pets" },
		{ Id: 8, Name:  "🏦 House Credit" },
		{ Id: 9, Name:  "🏝 Leisure General" },
		{ Id: 2, Name:  "✈️ Leisure Tourism" },
		{ Id: 10, Name: "⚽️ Leisure Hobby" },
		{ Id: 11, Name: "🎁 Leisure Presents" },
		{ Id: 12, Name: "📽 Leisure Entertainment" },
		{ Id: 16, Name: "👕 Clothing" },
		{ Id: 19, Name: "🚑 Health" },
		{ Id: 17, Name: "🪣 Other" },
		{ Id: 13, Name: "🏆 Personal" },
		{ Id: 3, Name:  "👨‍💻 Work" },
		{ Id: 18, Name: "🧹 Corrections" },
		{ Id: 20, Name: "❔ Unknown" },
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
