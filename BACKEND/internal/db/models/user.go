package models

import (
	"gorm.io/gorm"
)

// swagger:model
type User struct {
	gorm.Model
	// The name of the User
	// required: true
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
