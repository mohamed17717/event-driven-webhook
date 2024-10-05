package models

import "time"

type Change struct {
	ID         uint      `gorm:"primaryKey"`
	UserID     uint      `gorm:"index"`
	ActionID   uint      `gorm:"index"`
	Data       string    `gorm:"type:json"` // json field
	Identifier string    `gorm:"size:128"`
	When       time.Time `gorm:"column:when"`
	CreatedAt  time.Time `gorm:"default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"default:current_timestamp"`
}
