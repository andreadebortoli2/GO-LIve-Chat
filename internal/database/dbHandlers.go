package database

import (
	"errors"
	"log"

	"github.com/andreadebortoli2/GO-Experiment-and-Learn/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 12

// Login return the user if exist
func Login(email, password string) (models.User, error) {
	var result models.User
	tx := dbConn.SQLite3.Table("users").Select("*").Where("email = ?", email).Scan(&result)
	err := tx.Error
	if err != nil {
		log.Println(err)
		return result, err
	}

	if result.Password == "" {
		return result, errors.New("user not registrered")
	}
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return result, err
	}
	return result, nil
}

// AddUser add new user to db
func AddUser(userName, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		log.Println(err)
		return err
	}

	newUser := models.User{
		UserName: userName,
		Email:    email,
		Password: string(hashedPassword),
	}

	result := dbConn.SQLite3.Create(&newUser)
	err = result.Error
	if err != nil {
		return err
	}

	return nil
}
