package render

import (
	"github.com/alexedwards/scs/v2"
	"github.com/kasuki2/bookings/internal/config"
	"github.com/kasuki2/bookings/internal/models"
	"testing"
	"encoding/gob"
	"time"
	"os"
	"net/http"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {

	// change this to true when in production
	testApp.InProduction = false

	// what is in the session? primitives are okay
	// other types have to be told to session

	gob.Register(models.Reservation{})

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp


	os.Exit(m.Run())
}


type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write( b []byte) (int, error) {
	length := len(b)
	return length, nil
}