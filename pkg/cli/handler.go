package cli

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/entities"
	"github.com/randaalex/finance_bot/pkg/firefly"
	"github.com/randaalex/finance_bot/pkg/telegram"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type Storage interface {
	CreateTransactionsLog(ctx context.Context, arg db.CreateTransactionsLogParams) (db.TransactionsLog, error)
	GetTransactionsLogByDescription(ctx context.Context, description string) (db.TransactionsLog, error)
}

func newLogger() *logrus.Logger {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	//log.SetReportCaller(true)

	return log
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
		Host:       settings.FireflyHost,
		Scheme:     settings.FireflyScheme,
		HTTPClient: tc,
		OperationServers: map[string]firefly.ServerConfigurations{
			"TransactionsApiService.StoreTransaction":  {{URL: ""}},
			"TransactionsApiService.GetTransaction":    {{URL: ""}},
			"TransactionsApiService.UpdateTransaction": {{URL: ""}},
		},
	}), nil
}

func newTelegramBot(storage Storage, logger *logrus.Logger, fireflyClient *firefly.APIClient, settings *entities.Settings) (*telegram.Bot, error) {
	return telegram.NewBot(storage, logger, fireflyClient, settings, getAccounts(), getCategories()), nil
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
			Id:            7,
			Name:          "AlfaBank N1 (BYN)",
			CurrencyCode:  "BYN",
			AccountNumber: "7858",
		},
		{
			Id:            8,
			Name:          "AlfaBank N1 (USD)",
			CurrencyCode:  "USD",
			AccountNumber: "6185",
		},
	}
}

func getCategories() *[]entities.Category {
	return &[]entities.Category{
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "⛽️ Car Fuel", Id: 14},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🛠 Car Maintenance", Id: 15},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "👕 Clothing", Id: 16},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "📽 Entertainment", Id: 12},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🍔 Food / Restaurants", Id: 5},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🛒 Groceries", Id: 4},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🚑 Health", Id: 19},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "⚽️ Hobby", Id: 10},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🏠 House Bills", Id: 1},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🏠 House Purchases", Id: 6},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🪣 Other", Id: 17},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🏆 Personal", Id: 13},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🐈 Pets", Id: 7},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🎁 Presents", Id: 11},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "✈️ Tourism", Id: 2},
		{Type: entities.TransactionSplitTypeWithdrawal, Name: "🚕 Transport", Id: 24},

		{Type: entities.TransactionSplitTypeDeposit, Name: "💰 Cashback", Id: 23},
		{Type: entities.TransactionSplitTypeDeposit, Name: "👨‍💻 Work", Id: 3},

		{Type: "other", Name: "🧹 Corrections", Id: 18},
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err) // TODO: fix panic
	}
}
