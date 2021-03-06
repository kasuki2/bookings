package main

import (
	"github.com/kasuki2/bookings/internal/handlers"
	"github.com/kasuki2/bookings/internal/config"
	"github.com/kasuki2/bookings/internal/render"
	"github.com/kasuki2/bookings/internal/models"
	"github.com/kasuki2/bookings/internal/helpers"
	"github.com/alexedwards/scs/v2"
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/gob"
	"os"
)
const portNumber = ":8089"
var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger


func main() {

		// change this to true when in production
		app.InProduction = false

		// what is in the session? primitives are okay
		// other types have to be told to session
	
		gob.Register(models.Reservation{})
	
		session = scs.New()
		session.Lifetime = 24 * time.Hour
		session.Cookie.Persist = true
		session.Cookie.SameSite = http.SameSiteLaxMode
		session.Cookie.Secure = app.InProduction
	
		app.Session = session
	
		tc, err := render.CreateTemplateCache()
		if err != nil {
			log.Fatal("cannot create template cache")
			//return err
		}
	
		app.TemplateCache = tc
		app.UseCache =false
	
		repo := handlers.NewRepo(&app)
		handlers.NewHandlers(repo)
	
		render.NewTemplates(&app)
			
	


	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)
	

	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	//http.ListenAndServe(portNumber, nil)

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
	


}

func run() error {

		

		// what is in the session? primitives are okay
		// other types have to be told to session
	
		gob.Register(models.Reservation{})

		// change this to true when in production
		app.InProduction = false

		infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
		app.InfoLog = infoLog

		errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
		app.ErrorLog = errorLog


		session = scs.New()
		session.Lifetime = 24 * time.Hour
		session.Cookie.Persist = true
		session.Cookie.SameSite = http.SameSiteLaxMode
		session.Cookie.Secure = app.InProduction
	
		app.Session = session
	
		tc, err := render.CreateTemplateCache()
		if err != nil {
			log.Fatal("cannot create template cache")
			return err
		}
	
		app.TemplateCache = tc
		app.UseCache =false
	
		repo := handlers.NewRepo(&app)
		handlers.NewHandlers(repo)
	
		render.NewTemplates(&app)
		helpers.NewHelpers(&app)

	return nil
}