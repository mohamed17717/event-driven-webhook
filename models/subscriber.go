package models

import "time"

type Subscriber struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"index"`
	WebhookLink string    `gorm:"size:512"`
	SecretToken string    `gorm:"size:512"`
	IsVerified  bool      `gorm:"default:true"`
	IsActive    bool      `gorm:"default:true"`
	CreatedAt   time.Time `gorm:"default:current_timestamp"`
	UpdatedAt   time.Time `gorm:"default:current_timestamp"`
}
