package main

import (
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"test_webapp/pkg/config"
	"test_webapp/pkg/handlers"
	"test_webapp/pkg/render"
	"test_webapp/pkg/server"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	srv := server.NewServer(portNumber, &app)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
