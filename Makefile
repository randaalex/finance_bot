BINARY=./bin/finance_bot
SRC=./main.go
APP_ENV?=development
FIREFLY_API_VERSION=1.4.0

ifneq (,$(wildcard ./app.env))
	include app.env
	export
endif

.PHONY: install build docker-build docker-push run sqlc firefly-api migrate migrate-down migrate-status

install:
	go get -v github.com/rubenv/sql-migrate/...
	echo "brew install sqlc"

build:
	go build -i -v -o $(BINARY) $(SRC)

docker-build:
	docker build . --tag ghcr.io/randaalex/finance_bot:latest

docker-push:
	docker push ghcr.io/randaalex/finance_bot:latest

run:
	go run main.go run

test:
	go test ./...

sqlc:
	cd config && sqlc generate && cd ..

firefly-api:
	rm -rf pkg/firefly
	openapi-generator generate -i https://api-docs.firefly-iii.org/firefly-iii-$(FIREFLY_API_VERSION).yaml -g go --package-name firefly -o pkg/firefly --additional-properties=generateInterfaces=true,enumClassPrefix=true
	# Configuration Model conflicts with global ApiClient Configuration
	#rm pkg/firefly/api_configuration.go
	echo "package firefly\n\ntype ConfigurationApi interface {}\ntype ConfigurationApiService service" > pkg/firefly/api_configuration.go
	rm -f pkg/firefly/model_configuration.go
	rm -f pkg/firefly/model_configuration_data.go
	rm -f pkg/firefly/model_configuration_update.go
	# Drop internal go.mod's
	rm -f pkg/firefly/git_push.sh
	rm -f pkg/firefly/go.mod
	rm -f pkg/firefly/go.sum

mocks:
	rm -rf pkg/mocks && mockery --all --keeptree --dir pkg --output pkg/mocks

migrate:
	sql-migrate up -env="${APP_ENV}" -config=config/dbconfig.yml

migrate-down:
	sql-migrate down -env="${APP_ENV}" -config=config/dbconfig.yml

migrate-status:
	sql-migrate status -env="${APP_ENV}" -config=config/dbconfig.yml
