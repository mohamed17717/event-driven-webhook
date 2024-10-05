package models

import "time"

type SubscriberAction struct {
	ID           uint      `gorm:"primaryKey"`
	SubscriberID uint      `gorm:"index"`
	ActionID     uint      `gorm:"index"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}
