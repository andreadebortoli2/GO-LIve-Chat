package database

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/andreadebortoli2/GO-Live-Chat/internal/models"
	"golang.org/x/crypto/bcrypt"
)

var bcryptCost = 12

// Login return the user if exist
func Login(email, password string) (models.User, error) {
	var result models.User
	tx := dbConn.SQLite3.Table("users").Select("*").Where("email = ?", email).Scan(&result)
	if err := tx.Error; err != nil {
		log.Println(err)
		return result, errors.New("cannot find the user into the database")
	}

	if tx.RowsAffected <= 0 {
		return result, errors.New("user not registrered")
	}

	err := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password))
	if err != nil {
		log.Println(err)
		return result, errors.New("the password is not correct")
	}
	return result, nil
}

// AddUser add new user to db
func AddUser(userName, email, password string) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
	if err != nil {
		log.Println(err)
		return errors.New("failed encrypting the password")
	}

	newUser := models.User{
		UserName: userName,
		Email:    email,
		Password: string(hashedPassword),
	}

	result := dbConn.SQLite3.Create(&newUser)
	err = result.Error
	if err != nil {
		es1 := strings.Split(err.Error(), ".")
		es2 := strings.Split(es1[1], " ")
		return fmt.Errorf("a user with this %s already exist", es2[0])
	}

	if result.RowsAffected <= 0 {
		return errors.New("database error, user not registrered")
	}

	return nil
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	// result := []models.User{}
	// tx := dbConn.SQLite3.Raw("SELECT * FROM users ORDER BY id")
	// GORM syntax to not get soft deleted users
	tx := dbConn.SQLite3.Find(&users).Order("id")
	if err := tx.Error; err != nil {
		log.Println(err)
		return users, errors.New("cannot find the users into the database")
	}
	if tx.RowsAffected <= 0 {
		return users, errors.New("users not found")
	}

	return users, nil
}

func SetModerator(accLvl string, id string) error {
	var user models.User
	tx := dbConn.SQLite3.First(&user, id)
	if err := tx.Error; err != nil {
		log.Println(err)
		return err
	}
	if tx.RowsAffected <= 0 {
		return errors.New("user not found")
	}

	var accLvlReverse string

	if accLvl == "1" {
		accLvlReverse = "2"
	} else {
		accLvlReverse = "1"
	}

	user.AccessLevel = accLvlReverse
	tx = dbConn.SQLite3.Save(&user)
	if err := tx.Error; err != nil {
		log.Println(err)
		return err
	}
	if tx.RowsAffected <= 0 {
		return errors.New("user not updated")
	}

	return nil
}

func DeleteUserByID(id string) error {
	tx := dbConn.SQLite3.Delete(&models.User{}, id)
	if err := tx.Error; err != nil {
		log.Println(err)
		return err
	}
	if tx.RowsAffected <= 0 {
		return errors.New("user not deleted")
	}
	return nil
}

func GetLastMessages() ([]models.Message, error) {
	var messages []models.Message

	tx := dbConn.SQLite3.Order("id DESC").Limit(5).Preload("User").Find(&messages)
	if err := tx.Error; err != nil {
		log.Println(err)
		return messages, err
	}
	if tx.RowsAffected <= 0 {
		return messages, errors.New("messages not found")
	}

	for i, j := 0, len(messages)-1; i < j; i, j = i+1, j-1 {
		messages[i], messages[j] = messages[j], messages[i]
	}

	return messages, nil
}

func GetOlderMessages() ([]models.Message, error) {
	var messages []models.Message

	tx := dbConn.SQLite3.Offset(5).Preload("User").Find(&messages)
	if err := tx.Error; err != nil {
		log.Println(err)
		return messages, err
	}
	if tx.RowsAffected <= 0 {
		return messages, errors.New("messages not found")
	}

	return messages, nil
}
