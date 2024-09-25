package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName    string `gorm:"unique"`
	Email       string `gorm:"unique"`
	Password    string `gorm:"type:varchar(60);unique"`
	AccessLevel string
	Messages    []Message
}

func (p *User) BeforeCreate(tx *gorm.DB) error {
	if p.AccessLevel == "" {
		p.AccessLevel = "1"
	}
	return nil
}

type Message struct {
	gorm.Model
	Content string
	UserID  uint64
	User    User
}
