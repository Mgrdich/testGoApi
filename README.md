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

## Before Opening PR 
* make sure to run `make lint` to check for linter bugs

## For Migrations
* if you want to create new migration run `make migrate-create n=name-of-migration` it will create new file in `internal/db/migrations` with the correct name edit the up and down statements.
* make you database up to date `make migrate-up`.
* with `make migrate-down` please remove the associated file with it as well.
* Make migrations file for `atomic` operations.


## For Linting on the local machine
Install `golangci-lint` [Link](https://golangci-lint.run/welcome/install/) recommended way is to put it in the golang binary directory

## Testing 
Everything should have corresponding test suites except the repository for now

```shell 
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin VERSION_NUMBER_HERE
```

add the `$(go env GOPATH)` to `PATH`

## For Swagger 
To generate Swagger documentation for the API, first, add annotations to the controllers. Here's an example:
```
// HandleCreateUser creates a new user.
// @Summary Create a new user
// @Description Creates a new user with the provided details.
// @Tags user
// @Param username query string true "Username of the new user"
// @Param email query string true "Email address of the new user"
// @Param password query string true "Password of the new user"
// @Success 201 {object} UserDTO "User created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request format"
// @Failure 500 {object} ErrorResponse "Internal server error"
// @Router /api/v1/users [post]
func HandleCreateUser(w http.ResponseWriter, r *http.Request) {
    // Handler logic to create a new user
}
```
Once the annotations are added, you can generate Swagger documentation by running:
```bash
make gen-swagger
```
After running the command, you can access the documentation of your API by running the project and visiting:
http://localhost:8080/swagger/index.html

### For VsCode users
```json
"go.lintTool" : "golangci-lint",
"go.lintFlags": [
"--fast"
]
```