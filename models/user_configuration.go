package models

type UserConfiguration struct {
	ID               uint `gorm:"primaryKey"`
	UserID           uint `gorm:"index"`
	LatestChangeOnly bool `gorm:"default:true"`
	RetryFailure     bool `gorm:"default:true"`
	MaxRetries       int  `gorm:"default:5"`
}
