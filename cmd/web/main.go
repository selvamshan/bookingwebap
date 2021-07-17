package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/selvamshan/bookingwebapp/pkg/config"
	"github.com/selvamshan/bookingwebapp/pkg/handlers"
	"github.com/selvamshan/bookingwebapp/pkg/render"
)

const portNumber = ":8080"
var app config.AppConfig
var session *scs.SessionManager


func main() {	

	// whether App in production or not change this true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24*time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc

	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandler(repo)

	render.NewTemplate(&app)

	routes(&app)

	fmt.Println(fmt.Sprintf("Starting app on port no %s", portNumber))

	srv := http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
