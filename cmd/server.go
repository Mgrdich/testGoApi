package main

import (
	"context"
	"log"

	"testGoApi.com/configs"
	"testGoApi.com/internal/db"
	"testGoApi.com/internal/routes"
	"testGoApi.com/internal/server"
)

func main() {
	ctx := context.Background()

	conn, err := db.ConnectPostgresql(configs.GetAppConfig().PostgresqlUrl)
	if err != nil {
		log.Fatalf("error in connection %v/n", err)
	}
	defer conn.Close(context.Background())

	apiServer := server.NewServer(conn)
	routes.AddRoutes(apiServer)
	apiServer.Start(ctx)
}
