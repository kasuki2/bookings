package main

import (
	"testing"

)

func TestRun(t *testing.T) {
	err := run()

	if err != nil {
		t.Error("filed run()")
	}
}

/*


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
	}

	app.TemplateCache = tc
	app.UseCache =false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

*/