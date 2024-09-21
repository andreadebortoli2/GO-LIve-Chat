package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm: "unique"`
	Email       string `gorm: "unique"`
	Password    string `gorm: "unique"`
	AccessLevel string
}

func (p *User) BeforeCreate(tx *gorm.DB) error {
	if p.AccessLevel == "" {
		p.AccessLevel = "1"
	}
	return nil
}
