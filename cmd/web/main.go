package main

import (
	"bookings-udemy/internal/config"
	"bookings-udemy/internal/driver"
	"bookings-udemy/internal/handlers"
	"bookings-udemy/internal/helpers"
	"bookings-udemy/internal/models"
	"bookings-udemy/internal/render"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

// main is the main function
func main() {

	db,err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	defer close(app.MailChan)

	fmt.Println("Starting mail listner")
	listenForMail()

	/*msg:=models.MailData{
		To: "john@do.ca",
		From: "ma@here.com",
		Subject:"Testing mail from Go server.",
		Content: "",
	}

	app.MailChan <-msg*/

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func run() (*driver.DB,error) {
	// What am I going to put in Session

	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan:=make(chan models.MailData)
	app.MailChan=mailChan

	// change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up the session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to database

	log.Println("Connecting to Database")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=password")
	if err!= nil{
		log.Fatal("Cannot connect to database ! Dying...")
	}
	log.Println("Connected to Database!")
	

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil,err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app,db)
	handlers.NewHandlers(repo)

	render.NewRenderer(&app)
	helpers.NewHelpers(&app)
	return db,nil

}
