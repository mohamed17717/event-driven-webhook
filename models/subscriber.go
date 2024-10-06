package models

import (
	"event-driven-webhook/config"
	"time"
)

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

func GetSubscribersForAction(actionID any) []Subscriber {
	var subscribers []Subscriber
	config.DB.Joins("JOIN subscriber_actions ON subscribers.id = subscriber_actions.subscriber_id").
		Where("subscriber_actions.action_id = ?", actionID).
		Find(&subscribers)

	return subscribers
}
