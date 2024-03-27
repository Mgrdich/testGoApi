include .env
export

generate:
	sqlc generate

diff:
	sqlc diff

sqlc-lint:
	sqlc vet

vet:
	go vet ./...

lint: sqlc-lint vet
	golangci-lint run

build_ci:
	CGO_ENABLED=0 GOOS=linux go build -o server cmd/server.go

build: generate build_ci

run: generate
	go run cmd/server.go

install:
	go mod download

migrate-create:
	goose create ${n} sql

migrate-up:
	goose up

migrate-up-to:
	goose up-to ${v}

migrate-down:
	goose down

migrate-down-to:
	goose down-to ${v}

migrate-status:
	goose status

migrate-version:
	goose version

migrate-validate:
	goose validate