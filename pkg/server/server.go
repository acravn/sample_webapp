package server

import (
	"log"
	"net/http"
	"sample_webapp/pkg/config"
	"sample_webapp/pkg/routes"
)

func NewServer(port string, app *config.AppConfig) *http.Server {
	log.Printf("Starting server on port %s...\n", port)
	return &http.Server{
		Addr:    port,
		Handler: routes.Routes(app),
	}
}
