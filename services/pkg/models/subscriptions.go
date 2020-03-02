package models

import "time"

//Subscription Model
type Subscription struct {
	ID        int        `gorm:"id"`
	UserID    int        `gorm:"user_id"`
	ServiceID int        `gorm:"service_id"`
	PackID    int        `gorm:"pack_id"`
	ChannelID int        `gorm:"channel_id"`
	Duration  uint       `gorm:"duration"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}

//SubscriptionDetails Model
type SubscriptionDetails struct {
	ID          int    `gorm:"id"`
	UserID      int    `gorm:"user_id"`
	ServiceName string `gorm:"service_name"`
	PackName    string `gorm:"pack_name"`
	ChannelName string `gorm:"channel_name"`
}
