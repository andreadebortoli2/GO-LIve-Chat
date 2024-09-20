package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm: "unique`
	Email       string `gorm: "unique`
	Password    string `gorm: "unique`
	AccessLevel int
}
