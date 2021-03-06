package main

import (
	"fmt"
	"github.com/Tomonaga007/bookings/internal/config"
	"github.com/Tomonaga007/bookings/internal/handlers"
	"github.com/Tomonaga007/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const addr = ":8000"
var session *scs.SessionManager

func main() {
	var app config.AppConfig
	app.InProduction = true

	session = scs.New()
	session.Lifetime = 24* time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil{
		log.Fatal("Cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	svr := http.Server{Addr: addr, Handler: routes(&app)}
	fmt.Println("starting server at: ",addr )
	log.Fatal(svr.ListenAndServe())

}
