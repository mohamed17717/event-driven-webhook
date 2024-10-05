package models

import "time"

type WebhookLog struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       uint   `gorm:"index"`
	ChangeID     uint   `gorm:"index"`
	SubscriberID uint   `gorm:"index"`
	Status       string `gorm:"size:32;default:'pending'"`
	Retries      int    `gorm:"default:0"`
	StatusCode   int
	ResponseText string    `gorm:"size:512"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}
