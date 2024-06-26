package server

import (
	"fmt"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/zk-ticketing/zk-ticketing/backend/jwt"
	"github.com/zk-ticketing/zk-ticketing/backend/openapi"
	"github.com/zk-ticketing/zk-ticketing/backend/repos"
	"github.com/zk-ticketing/zk-ticketing/backend/service"
)

type Server struct {
	dbClient   *repos.Client
	port       int
	apiService *service.APIService
	jwtService *jwt.Service
}

func New(
	dbClient *repos.Client,
	port int,
	apiService *service.APIService,
	jwtService *jwt.Service,
) *Server {
	return &Server{
		dbClient:   dbClient,
		port:       port,
		apiService: apiService,
		jwtService: jwtService,
	}
}

func (s *Server) Start() {
	defaultAPIController := openapi.NewDefaultAPIController(s.apiService)
	defaultRouter := openapi.NewRouter(defaultAPIController)
	router := authMiddleware(defaultRouter, s.jwtService)
	router = corsMiddleware(router)

	log.Printf("Starting server on port %d", s.port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", s.port), router)
	if err != nil {
		log.Fatal().Msgf("Server failed with error: %v", err)
	}
}
