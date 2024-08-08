package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/simplifywoopii88/bookings/pkg/config"
	"github.com/simplifywoopii88/bookings/pkg/handlers"
	"github.com/simplifywoopii88/bookings/pkg/render"
)

const portNumber = ":8080"



var app config.AppConfig
var session *scs.SessionManager
func main(){
	

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create temlate cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)
	 
	// http.HandleFunc("/", handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	fmt.Printf("listening server http://localhost%s\n", portNumber)
	// if err := http.ListenAndServe(portNumber, nil); err != nil {
	// 	log.Fatal(err)
	// }

	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}