package database

import (
	"log"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var user *models.User

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("GO_exp_learn.db"), &gorm.Config{})
	if err != nil {
		return db, err
	}
	db.Exec("DROP TABLE users")
	execMigrations(db)
	userSeeder(db)
	return db, nil
}

// execMigrations execute all the migrations
func execMigrations(db *gorm.DB) {
	// add all models's structs to AutoMigrate
	db.AutoMigrate(&user)
}

func userSeeder(db *gorm.DB) {
	// TODO! hash passwords
	users := []*models.User{
		{UserName: "admin", Email: "admin@admin.com", Password: "password", AccessLevel: 2},
		{UserName: "moderator", Email: "moderator@moderator.com", Password: "password", AccessLevel: 1},
		{UserName: "user", Email: "user@user.com", Password: "password"},
	}

	result := db.Create(users)
	if err := result.Error; err != nil {
		log.Println(err)
	}
	log.Printf("added %d initial users", result.RowsAffected)
}
