package database

import (
	"log"

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

	// all passwords are: password
	users := []*models.User{
		{
			UserName:    "admin",
			Email:       "admin@admin.com",
			Password:    "$2a$12$39JTEON1eLjhQ4uHq/L8SuQNn9kUgqCuCA3LmSZ3A9iJK6Ay82VvC",
			AccessLevel: "3",
		},
		{
			UserName:    "moderator",
			Email:       "moderator@moderator.com",
			Password:    "$2a$12$yMUW6GklJCw3ehtbs9kDQ.AtlTYPCLnimGNgWN6BH9bjvAOlXge1G",
			AccessLevel: "2",
		},
		{
			UserName: "user",
			Email:    "user@user.com",
			Password: "$2a$12$JCdNB2z/3YwQhUjd1TVlDeaf4ULeNmNoKcj1V6qWUUFKjkC7b.q2q",
		},
	}

	result := db.Create(users)
	if err := result.Error; err != nil {
		log.Println(err)
	}
	log.Printf("added %d initial users", result.RowsAffected)
}
