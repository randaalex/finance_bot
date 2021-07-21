package cli

import (
	"context"
	"database/sql"
	"github.com/randaalex/finance_bot/pkg/bot"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"github.com/randaalex/finance_bot/pkg/processors/fileprocessor"
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
		Host: settings.FireflyBaseUrl,
		HTTPClient: tc,
	}), nil
}

func newCategorySelector() (fileprocessor.CategorySelectorInterface, error) {
	categories := []fileprocessor.Category{
		{ID: 3, Name: "Work"},
		{ID: 4, Name: "🛒 Food Groceries"},
		{ID: 5, Name: "🍔 Food Restaurants"},
		{ID: 1, Name: "House Bills"},
		{ID: 6, Name: "House Purchases"},
		{ID: 7, Name: "House Pets"},
		{ID: 8, Name: "House Credit"},
		{ID: 9, Name: "Leisure General"},
		{ID: 2, Name: "Leisure Tourism"},
		{ID: 10, Name: "Leisure Hobby"},
		{ID: 11, Name: "Leisure Presents"},
		{ID: 12, Name: "Leisure Entertainment"},
		{ID: 13, Name: "Personal"},
		{ID: 19, Name: "Health"},
		{ID: 14, Name: "Car Fuel"},
		{ID: 15, Name: "Car Maintenance"},
		{ID: 16, Name: "Clothing"},
		{ID: 17, Name: "Other"},
		{ID: 18, Name: "Corrections"},
	}

	return fileprocessor.NewCategorySelector(&categories), nil
}

func newBot(storage Storage, fireflyClient *firefly.APIClient, settings *entities.Settings) (*bot.Bot, error) {
	return bot.NewBot(storage, fireflyClient, settings), nil
}

func getSettings() *entities.Settings {
	var settings entities.Settings

	_ = viper.Unmarshal(&settings)

	return &settings
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
