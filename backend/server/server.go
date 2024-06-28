package server

import (
	"fmt"
	"net/http"

	"github.com/proof-pass/proof-pass/backend/jwt"
	"github.com/proof-pass/proof-pass/backend/openapi"
	"github.com/proof-pass/proof-pass/backend/repos"
	"github.com/proof-pass/proof-pass/backend/service"
	"github.com/rs/zerolog/log"
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
