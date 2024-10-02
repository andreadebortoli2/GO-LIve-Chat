package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/config"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/database"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/handlers"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/models"
	"github.com/andreadebortoli2/GO-Live-Chat/internal/render"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
)

var appConfig config.AppConfig

var session *sessions.CookieStore

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Connected to database")
	sqlDB, _ := db.SQLite3.DB()
	defer func() {
		sqlDB.Close()
	}()

	// set the session
	gob.Register(models.User{})
	session = sessions.NewCookieStore([]byte(os.Getenv("MY_SECRET_KEY")))
	appConfig.Session = session

	repo := handlers.NewRepo(&appConfig, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&appConfig)

	router := Router(&appConfig)

	fmt.Println("serving on port 8080")
	_ = http.ListenAndServe(":8080", router)

}
