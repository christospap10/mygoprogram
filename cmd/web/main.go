package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/christospap10/mygoprogram/pkg/config"
	"github.com/christospap10/mygoprogram/pkg/handlers"
	"github.com/christospap10/mygoprogram/pkg/render"
)

// Port
const portNumber = ":8080"

// This is the main function
var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error creating template cache: ", err)
	}

	app.TemplateCache = tc
	app.UseCache = false // Development mode
	// Set the handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Server is running on port %s", portNumber))
	// Listen on port 8080, IP defaults to
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
