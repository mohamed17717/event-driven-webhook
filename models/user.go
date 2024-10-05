package models

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:50;unique"`
	Email     string    `gorm:"size:100;unique"`
	Password  string    `gorm:"size:255"`
	FirstName string    `gorm:"size:50"`
	LastName  string    `gorm:"size:50"`
	AuthKey   string    `gorm:"size:512;unique"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
