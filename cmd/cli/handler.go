package cli

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/randaalex/finance_bot/pkg/db"
)

type Handler struct {
	DB     db.DBTX
	Parse   *ParseHandler
}

func NewHandler() *Handler {
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

	return &Handler{
		DB:    dbConnection,
		Parse: newParseHandler(dbStorage),
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
