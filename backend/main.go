package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/NFTGalaxy/zk-ticketing-server/openapi"
	"github.com/NFTGalaxy/zk-ticketing-server/repos"
	"github.com/NFTGalaxy/zk-ticketing-server/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	RestPort         int    `default:"3000"`
	PostgresHost     string `default:"postgres.default.svc.cluster.local"` // TODO: namespace
	PostgresDatabase string `default:"db"`
	PostgresPort     int    `default:"5432"`
	PostgresUsername string `required:"true"`
	PostgresPassword string `required:"true"`
}

func main() {
	// load the configuration from the environment variables
	var cfg config
	envconfig.MustProcess("backend", &cfg)

	// connect to the database
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDatabase))
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()
	dbClient := repos.NewClient(pool)

	// create the API service
	apiService := service.NewAPIService(dbClient)

	defaultAPIController := openapi.NewDefaultAPIController(apiService)
	router := openapi.NewRouter(defaultAPIController)

	log.Printf("Starting server on port %d", cfg.RestPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.RestPort), router))
}
