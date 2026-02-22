package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       int
	Username string `gorm:"unique"`
	Password string
	Balance  int
}

type CheckIn struct {
	gorm.Model
	ID          int
	UserID      int
	OrderBookID int
	CheckInAt   time.Time
	CheckOutAt  time.Time
	Status      string
}
