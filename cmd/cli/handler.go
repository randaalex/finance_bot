package cli

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/randaalex/finance_bot/pkg/processors/fileprocessor"

	_ "github.com/lib/pq"
	"golang.org/x/oauth2"

	"github.com/randaalex/finance_bot/pkg/db"
	"github.com/randaalex/finance_bot/pkg/firefly"
)

type Storage interface {
	CreateMapping(ctx context.Context, arg db.CreateMappingParams) (db.Mapping, error)
	GetMappingByTransactionDetails(ctx context.Context, transactionDetails string) (db.Mapping, error)

	CreateProcessedTransaction(ctx context.Context, arg db.CreateProcessedTransactionParams) (db.ProcessedTransaction, error)
	GetProcessedTransactionByFireflyID(ctx context.Context, fireflyID sql.NullInt32) (db.ProcessedTransaction, error)
	GetProcessedTransactionByHash(ctx context.Context, hash string) (db.ProcessedTransaction, error)
}

func newDBClient() (Storage, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "postgres"
		dbname   = "finance_bot_development"
	)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	dbConnection, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	err = dbConnection.Ping()
	CheckError(err)

	fmt.Println("Connected!")

	dbStorage := db.New(dbConnection)

	return dbStorage, nil
}

func newFireflyClient() (*firefly.Client, error) {
	token := "eyJ0eXAiOiJKV1QiLCJhbGciOiJSUzI1NiJ9.eyJhdWQiOiI1IiwianRpIjoiN2ExY2Q2OTg5NmEzMWRiNWY2YTdmOGYxOTMzMzBiMzhjMjZjZjYwYThiNjNmMGRlMGYwNmU2NWJhMmU3ZmYxMTYyYTM5YTgxYjJjMGQ5NWYiLCJpYXQiOjE2MTY3NDk4NTksIm5iZiI6MTYxNjc0OTg1OSwiZXhwIjoxNjQ4Mjg1ODU5LCJzdWIiOiIxIiwic2NvcGVzIjpbXX0.MUNgAHR9UHcWMk7LB30JEfqOFc6zC46atvZzLGZiMpTqcCT0AOSSz258EPPtuLRCPdion1MBwOsV4QpHElMvGruEfFR-PnDU5PDBZ77L0iHKE-asOQ5jjID5ARckkB3H25W7C95PQuXDKtmrS8irtxTVrWx17Ad9krvYIxEIjpvoNGxPMMnH9S9uuyR5NISfJbqdScvP33EFptKflSk_uGXbK0b9RFjYbjWDSlbPC8Ux7DJyaFljvwMIz9kIbn4TkL9u9868RUUHHHL8SPVOFTKF_BOQ2qmUzZb8nibXybAEVoQza4rHc8srBf1iFjY0h9uNMYEvpYBdKWFHk5GTKCKPpz__iYKIAEp1uJBSGYU15q1lAz4u1a6-0Bym2JnNV5lEYM42pFsVn8gHOIvnpdRAfVi7OI-ZSxUt9EdRU_7bGoY_kqxSh_XQ3I7IzMW2xa1v-YzH8vGKleFN5J7ZPyVZQ945pxvC1olyHFFO0pBebb3Iy1pDrjjAnQcWmvvz0JLudlbukIpOvb6M6QPxCcQNqX99Wh_xSYZUq3uQw3vyK_pK1AhIUgbShSm36GUEDZzqS90vym__RRvBb9xJTjFCZnNKiX8WD5oJJE2JMlWLij9uAEPUa_haljs9ujcqK8VXKDy1aius6f6ZXpG-pJ46UW8V-3_lw5jK-rQt4B0"
	baseUrl := "http://192.168.0.2:8081/api/v1/"

	ctx := context.TODO()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return firefly.NewClient(tc, baseUrl), nil
}

func newCategorySelector() (fileprocessor.CategorySelectorInterface, error) {
	categories := []fileprocessor.Category{
		{ID: 3, Name: "Work"},
		{ID: 4, Name: "Food Grocering"},
		{ID: 5, Name: "Food Restaurants"},
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

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
