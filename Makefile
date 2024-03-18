include .env
export

generate:
	sqlc generate

diff:
	sqlc diff

lint:
	sqlc vet

build: generate
	CGO_ENABLED=0 GOOS=linux go build -o server cmd/server.go

run: generate
	go run cmd/server.go

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