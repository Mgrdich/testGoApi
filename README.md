# Minor API project for Golang

In order to run the project you need to have `sqlc` please check https://github.com/sqlc-dev/sqlc

Project demonstrates how to create Go Backend project using
* `chi` for Router
* `sqlc` for compiling `sql` into type safe code
* `migrate` to handle the database migrations

## Start the application

create `.env` file
```shell
PORT=****
postgresql=postgres://username:password@localhost:5432/db-name
```

```shell
make run
```
