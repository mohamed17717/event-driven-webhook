package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey"`
	Username  string    `gorm:"size:50;unique"`
	Email     string    `gorm:"size:100;unique"`
	Password  string    `gorm:"size:255"`
	FirstName string    `gorm:"size:50"`
	LastName  string    `gorm:"size:50"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
}

func (user *User) AfterCreate(tx *gorm.DB) (err error) {
	// Create the profile after the user is created
	configuration := UserConfiguration{
		UserID: user.ID,
	}
	err = tx.Create(&configuration).Error
	return err
}
