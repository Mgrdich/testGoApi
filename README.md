# Minor API project for Golang

* In order to run the project you need to have `sqlc` please check https://github.com/sqlc-dev/sqlc
* Goose to handle Database migrations https://github.com/pressly/goose

Project demonstrates how to create Go Backend project using
* `chi` for Router
* `sqlc` for compiling `sql` into type safe code
* `goose` to handle the database migrations

## Start the application

create `.env` file at root
```shell
PORT=8080
POSTGRESQL=postgres://username:password@localhost:5432/db-name
GOOSE_DRIVER=postgres
GOOSE_DBSTRING="user=username password=password dbname=db-name sslmode=disable"
GOOSE_MIGRATION_DIR=./internal/db/migrations
```

```shell
make run
```