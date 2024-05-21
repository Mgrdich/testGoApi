package main

import (
	"context"
	"log"
	"testGoApi/internal/repository"

	"testGoApi/configs"
	"testGoApi/internal/db"
	"testGoApi/internal/routes"
	"testGoApi/internal/server"
	"testGoApi/internal/services"
)

func main() {
	ctx := context.Background()

	conn, pQueries, err := db.ConnectPostgresql(configs.GetAppConfig().PostgresqlUrl)
	if err != nil {
		log.Fatalf("error in connection %v/n", err)
	}
	defer conn.Close(context.Background())

	apiServer := server.NewServer(conn)
	routes.AddRoutes(apiServer, &routes.ApplicationServices{
		MovieService:  services.NewMoviesServiceImpl(repository.NewMoviesRepositoryImpl(pQueries)),
		PersonService: services.NewPersonServiceImpl(repository.NewPersonRepositoryImpl(pQueries)),
	})
	apiServer.Start(ctx)
}
