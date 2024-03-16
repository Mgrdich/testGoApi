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