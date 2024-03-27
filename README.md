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
GOOSE_DBSTRING=${POSTGRESQL}
GOOSE_MIGRATION_DIR=./internal/db/migrations
```

```shell
make run
```

## Before Opening MR 
* make sure to sure to run `make lint` to check for linter bugs

## For Migrations
* if you want to create new migration run `make migrate-create n=name-of-migration` it will create new file in `internal/db/migrations` with the correct name edit the up and down statements.
* make you database up to date `make migrate-up`.
* with `make migrate-down` please remove the associated file with it as well.
* Make migrations file for `atomic` operations.


## For Linting on the local machine
Install `golangci-lint` [Link](https://golangci-lint.run/welcome/install/) recommended way is to put it in the golang binary directory

```shell 
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin VERSION_NUMBER_HERE
```

add the `$(go env GOPATH)` to `PATH`


### For VsCode users
```json
"go.lintTool" : "golangci-lint",
"go.lintFlags": [
"--fast"
]
```