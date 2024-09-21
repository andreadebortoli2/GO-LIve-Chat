package database

import (
	"log"
	"net/http"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/config"
	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type DB struct {
	SQLite3   *gorm.DB
	AppConfig *config.AppConfig
}

var dbConn = &DB{}

var user *models.User

func ConnectDB() (*DB, error) {
	db, err := gorm.Open(sqlite.Open("GO_exp_learn.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	dbConn.SQLite3 = db
	db.Exec("DROP TABLE users")
	execMigrations(db)
	userSeeder(db)
	return dbConn, nil
}

// execMigrations execute all the migrations
func execMigrations(db *gorm.DB) {
	// add all models's structs to AutoMigrate
	db.AutoMigrate(&user)
}

func userSeeder(db *gorm.DB) {
	// TODO! hash passwords
	users := []*models.User{
		{UserName: "admin", Email: "admin@admin.com", Password: "password", AccessLevel: "3"},
		{UserName: "moderator", Email: "moderator@moderator.com", Password: "password", AccessLevel: "2"},
		{UserName: "user", Email: "user@user.com", Password: "password"},
	}

	result := db.Create(users)
	if err := result.Error; err != nil {
		log.Println(err)
	}
	log.Printf("added %d initial users", result.RowsAffected)
}

// Login return the user if exist
func Login(email, password string, r *http.Request) (models.User, error) {
	var result models.User
	err := dbConn.SQLite3.Table("users").Select("*").Where("email = ? AND password = ?", email, password).Scan(&result).Error
	if err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}
