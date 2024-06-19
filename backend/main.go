package main

import (
	"context"
	"fmt"

	"github.com/NFTGalaxy/zk-ticketing-server/jwt"
	"github.com/NFTGalaxy/zk-ticketing-server/repos"
	"github.com/NFTGalaxy/zk-ticketing-server/server"
	"github.com/NFTGalaxy/zk-ticketing-server/service"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog/log"
)

type config struct {
	RestPort         int    `default:"3000"`
	PostgresHost     string `default:"postgres.default.svc.cluster.local"` // TODO: namespace
	PostgresDatabase string `default:"db"`
	PostgresPort     int    `default:"5432"`
	PostgresUsername string `required:"true"`
	PostgresPassword string `required:"true"`
	JWTSecretKey     string `required:"true"`
	JWTExpiresSec    int64  `required:"true"`
}

func main() {
	// load the configuration from the environment variables
	var cfg config
	envconfig.MustProcess("backend", &cfg)

	// connect to the database
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDatabase))
	if err != nil {
		log.Fatal().Msgf("Unable to connect to database: %v", err)
	}
	defer pool.Close()
	dbClient := repos.NewClient(pool)

	// initialize JWT service
	jwtService := jwt.NewService(cfg.JWTSecretKey, cfg.JWTExpiresSec)

	// initialize API service
	apiService := service.NewAPIService(dbClient, jwtService)

	// create server
	server := server.New(dbClient, cfg.RestPort, apiService, jwtService)
	server.Start()
}
