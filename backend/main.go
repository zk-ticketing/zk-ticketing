package main

import (
	"context"
	"fmt"

	"github.com/NFTGalaxy/zk-ticketing-server/jwt"
	"github.com/NFTGalaxy/zk-ticketing-server/repos"
	"github.com/NFTGalaxy/zk-ticketing-server/server"
	"github.com/NFTGalaxy/zk-ticketing-server/service"
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type config struct {
	RestPort         int    `default:"3000"`
	PostgresHost     string `default:"postgres.app.svc.cluster.local"`
	PostgresDatabase string `default:"db"`
	PostgresPort     int    `default:"5432"`
	PostgresUsername string `required:"true"`
	PostgresPassword string `required:"true"`
	RedisAddr        string `default:"redis-master:6379"`
	JWTSecretKey     string `required:"true"`
	JWTExpiresSec    int64  `required:"true"`
}

func main() {
	// load the configuration from the environment variables
	var cfg config
	envconfig.MustProcess("backend", &cfg)

	// set up logger
	log.Logger = log.With().Caller().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	zerolog.DefaultContextLogger = &log.Logger

	// connect to the database
	pool, err := pgxpool.New(context.Background(), fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresDatabase))
	if err != nil {
		log.Fatal().Msgf("Unable to connect to database: %v", err)
	}
	defer pool.Close()
	dbClient := repos.NewClient(pool)

	// connect to Redis
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{Addrs: []string{cfg.RedisAddr}})

	// initialize JWT service
	jwtService := jwt.NewService(cfg.JWTSecretKey, cfg.JWTExpiresSec)

	// initialize API service
	apiService := service.NewAPIService(
		dbClient,
		redisClient,
		jwtService,
	)

	// create server
	server := server.New(dbClient, cfg.RestPort, apiService, jwtService)
	server.Start()
}
