package db

import (
	".com/internal/db/sqlc"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"log"
)

var pQueries *db.Queries

var postgresqlDbConnection *pgx.Conn

func GetPQueries() *db.Queries {
	return pQueries
}

func ConnectPostgresql(url string) (*pgx.Conn, error) {
	if postgresqlDbConnection != nil {
		return nil, errors.New("database instance exists")
	}

	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	postgresqlDbConnection = conn

	pQueries = db.New(conn)

	return conn, nil
}
