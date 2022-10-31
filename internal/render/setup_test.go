package render

import (
	"encoding/gob"
	"github.com/alexedwards/scs/v2"
	"Bookme/internal/config"
	"Bookme/internal/models"
	"log"
	"net/http"
	"os"
	"testing"
	"time"
)

var session *scs.SessionManager
var testApp config.AppConfig

//inital package
func TestMain(m *testing.M) {

	// what am I going to put in the session
	gob.Register(models.Reservation{})

	// change this to true when in production
	testApp.InProduction = false

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	testApp.InfoLog = infoLog

	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	testApp.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session
	//assing the testapp to app so the the app in render consist
	//the testAPp insetead of the app config
	app = &testApp

	os.Exit(m.Run())
}

type myWriter struct{}


//http response takes  is an interface contains Header Writer  and WriteHeader
//create atype that statidfy requirment so you could use it inseead http response
//writer 
func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {

}

func (tw *myWriter) Write(b []byte) (int, error) {
	//find the slice of the length of bytes 
	length := len(b)
	return length, nil
}