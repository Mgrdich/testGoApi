# Minor API project for Golang

* In order to run the project you need to have `sqlc` please check https://github.com/sqlc-dev/sqlc
* Download Atlas for Backend migrations https://atlasgo.io/getting-started

Project demonstrates how to create Go Backend project using
* `chi` for Router
* `sqlc` for compiling `sql` into type safe code
* `migrate` to handle the database migrations

## Start the application

create `.env` file at root
```shell
PORT=8080
POSTGRESQL=postgres://username:password@localhost:5432/db-name
```

```shell
make run
```

## Run a migrations 
Edit the schema file to the new structure then run
* `make migrate-create name=migration_name`