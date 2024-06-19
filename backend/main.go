package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/NFTGalaxy/zk-ticketing-server/openapi"
	"github.com/NFTGalaxy/zk-ticketing-server/service"
	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port int `default:"3000"`
}

func main() {
	var cfg config
	envconfig.MustProcess("backend", &cfg)

	apiService := service.NewAPIService()
	defaultAPIController := openapi.NewDefaultAPIController(apiService)

	router := openapi.NewRouter(defaultAPIController)

	log.Printf("Starting server on port %d", cfg.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), router))
}
