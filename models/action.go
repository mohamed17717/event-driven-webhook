package models

import "time"

type Action struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index"`
	EventName string    `gorm:"size:255"`
	CreatedAt time.Time `gorm:"default:current_timestamp"`
	UpdatedAt time.Time `gorm:"default:current_timestamp"`
}
