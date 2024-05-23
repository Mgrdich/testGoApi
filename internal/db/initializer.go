package db

import (
	"context"
	"errors"
	"log"

	"github.com/jackc/pgx/v5"
	db "testGoApi/internal/db/sqlc"
)

var pQueries *db.Queries

var postgresqlDbConnection *pgx.Conn

func GetPQueries() *db.Queries {
	if pQueries == nil {
		log.Fatalf("queries is not defined")
	}

	return pQueries
}

func ConnectPostgresql(url string) (*pgx.Conn, *db.Queries, error) {
	if postgresqlDbConnection != nil {
		return nil, nil, errors.New("database instance exists")
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	postgresqlDbConnection = conn

	pQueries = db.New(conn)

	return conn, pQueries, nil
}
