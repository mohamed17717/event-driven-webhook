package models

import (
	"gorm.io/gorm"
	"time"
)

type SubscriberAction struct {
	ID           uint      `gorm:"primaryKey"`
	SubscriberID uint      `gorm:"index"`
	ActionID     uint      `gorm:"index"`
	CreatedAt    time.Time `gorm:"default:current_timestamp"`
	UpdatedAt    time.Time `gorm:"default:current_timestamp"`
}

func CreateBulkActionsForOneSubscriber(db *gorm.DB, subscriberID uint, actionIDs []uint) error {
	var subscriberActions []SubscriberAction

	// Create a slice of SubscriberAction
	for _, actionID := range actionIDs {
		subscriberActions = append(subscriberActions, SubscriberAction{
			SubscriberID: subscriberID,
			ActionID:     actionID,
		})
	}

	// Perform the bulk insert
	if err := db.Create(&subscriberActions).Error; err != nil {
		return err
	}

	return nil
}

func CreateBulkSubscribersForOneAction(db *gorm.DB, actionID uint, subscriberIDs []uint) error {
	var subscriberActions []SubscriberAction

	// Create a slice of SubscriberAction
	for _, subscriberID := range subscriberIDs {
		subscriberActions = append(subscriberActions, SubscriberAction{
			SubscriberID: subscriberID,
			ActionID:     actionID,
		})
	}

	// Perform the bulk insert
	if err := db.Create(&subscriberActions).Error; err != nil {
		return err
	}

	return nil
}
