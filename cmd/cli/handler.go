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
	CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error)
	GetMappingByTransactionDetails(ctx context.Context, transactionDetails string) (db.Mapping, error)

	CreateProcessedTransaction(ctx context.Context, arg db.CreateProcessedTransactionParams) (db.ProcessedTransaction, error)
	GetProcessedTransactionByFireflyID(ctx context.Context, fireflyID sql.NullInt32) (db.ProcessedTransaction, error)
	GetProcessedTransactionByHash(ctx context.Context, hash string) (db.ProcessedTransaction, error)
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

func getAccounts() *map[int]string {
	// TODO: load from API
	return &map[int]string{
		7: "AlfaBank N1 (BYN)",
		8: "AlfaBank Card (USD)",
	}
}

func getCategories() *map[int]string {
	// TODO: load from API
	return &map[int]string{
		3:  "👨‍💻 Work",
		4:  "🛒 Food Groceries",
		5:  "🍔 Food Restaurants",
		1:  "🏠 House Bills",
		6:  "🏠 House Purchases",
		7:  "🐈 House Pets",
		8:  "🏦 House Credit",
		9:  "🏝 Leisure General",
		2:  "✈️ Leisure Tourism",
		10: "⚽️ Leisure Hobby",
		11: "🎁 Leisure Presents",
		12: "📽 Leisure Entertainment",
		13: "🏆 Personal",
		19: "🚑 Health",
		14: "⛽️ Car Fuel",
		15: "🛠 Car Maintenance",
		16: "👕 Clothing",
		17: " Other",
		18: " Corrections",
		20: "❔ Unknown",
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
