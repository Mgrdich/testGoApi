include .env
export

generate:
	sqlc generate

gen-swagger:
	swag init -g cmd/server.go --outputTypes go,yaml

diff:
	sqlc diff

sqlc-lint:
	sqlc vet

vet:
	go vet ./...

lint: sqlc-lint vet
	golangci-lint run

test:
	go test -v ./...

build_ci:
	CGO_ENABLED=0 GOOS=linux go build -o server cmd/server.go

build: generate gen-swagger build_ci

run: generate gen-swagger
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