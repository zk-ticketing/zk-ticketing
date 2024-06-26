package main

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zk-ticketing/zk-ticketing/backend/jwt"
	"github.com/zk-ticketing/zk-ticketing/backend/repos"
	"github.com/zk-ticketing/zk-ticketing/backend/server"
	"github.com/zk-ticketing/zk-ticketing/backend/service"
	"github.com/zk-ticketing/zk-ticketing/issuer/api/go/issuer/v1"
	"google.golang.org/grpc"
)

type config struct {
	RestPort                  int    `default:"3000"`
	PostgresHost              string `default:"postgres.app.svc.cluster.local"`
	PostgresDatabase          string `default:"db"`
	PostgresPort              int    `default:"5432"`
	PostgresUsername          string `required:"true"`
	PostgresPassword          string `required:"true"`
	RedisAddr                 string `default:"redis-master:6379"`
	IssuerAddr                string `default:"issuer.app.svc.cluster.local:9090"`
	IssuerChainID             int64  `default:"1"`   // TODO: change to actual context ID and set to requried
	EmailCredentialContextID  int64  `default:"111"` // TODO: change to actual context ID and set to requried
	TicketCredentialContextID int64  `default:"222"` // TODO: change to actual context ID and set to requried
	JWTSecretKey              string `required:"true"`
	JWTExpiresSec             int64  `required:"true"`
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

	// initialize issuer client at issuer.app.svc.cluster.local:9090
	issuerConn, err := grpc.NewClient(cfg.IssuerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal().Msgf("Unable to connect to issuer: %v", err)
	}
	issuerClient := issuer.NewIssuerServiceClient(issuerConn)

	// initialize API service
	apiService := service.NewAPIService(
		cfg.IssuerChainID,
		cfg.EmailCredentialContextID,
		cfg.TicketCredentialContextID,
		dbClient,
		redisClient,
		jwtService,
		issuerClient,
	)

	// create server
	server := server.New(dbClient, cfg.RestPort, apiService, jwtService)
	server.Start()
}
