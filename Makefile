BINARY=./bin/finance_bot
SRC=./main.go
APP_ENV?=development

.PHONY: build run sqlc migrate migrate-down migrate-status

build:
	go build -i -v -o $(BINARY) $(SRC)

run:
	go run main.go

sqlc:
	sqlc generate

mocks:
	rm -rf pkg/mocks && mockery --all --keeptree --dir pkg --output pkg/mocks

migrate:
	sql-migrate up -env="${APP_ENV}"

migrate-down:
	sql-migrate down -env="${APP_ENV}"

migrate-status:
	sql-migrate status -env="${APP_ENV}"
