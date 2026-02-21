package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Username string `gorm:"unique"`
	Password string
}
