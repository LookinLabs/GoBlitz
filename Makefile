_DEFAULT_GOAL := run

include .env

build:
	go build -o bin/main main.go

run:
	go run main.go

fumpt:
	gofumpt -w .

mod-vendor:
	go mod vendor

linter:
	@golangci-lint run

gosec:
	@gosec -quiet ./...

test:
	@env $(cat .env | xargs) go test -v ./tests/ -p 32

validate: linter gosec test

migrate-create:
	@goose -dir=migrations create "$(name)" sql

migrate-up:
	@goose -dir=migrations postgres "host=${PSQL_HOST} user=${PSQL_USERNAME} password=${PSQL_PASSWORD} dbname=${PSQL_DB} sslmode=disable" up

migrate-down:
	@goose -dir=migrations postgres "host=${PSQL_HOST} user=${PSQL_USERNAME} password=${PSQL_PASSWORD} dbname=${PSQL_DB} sslmode=disable" down
