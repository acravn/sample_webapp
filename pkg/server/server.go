package server

import (
	"net/http"
	"test_webapp/pkg/config"
	"test_webapp/pkg/routes"
)

func NewServer(port string, app *config.AppConfig) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: routes.Routes(app),
	}
}
