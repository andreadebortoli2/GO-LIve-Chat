package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/config"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/database"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/handlers"
)

var appConfig config.AppConfig
var session *scs.SessionManager

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected to database")
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// set the session parameters
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	appConfig.Session = session

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	router := Router()

	fmt.Println("serving on port 8080")
	_ = http.ListenAndServe(":8080", router)

}
