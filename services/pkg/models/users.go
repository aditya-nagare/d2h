package models

import "time"

//User Model
type User struct {
	ID        int        `gorm:"id"`
	Name      string     `gorm:"name"`
	Balance   uint       `gorm:"balance"`
	Email     string     `gorm:"email"`
	Phone     string     `gorm:"phone"`
	CreatedAt time.Time  `gorm:"created_at"`
	UpdatedAt time.Time  `gorm:"updated_at"`
	DeletedAt *time.Time `gorm:"deleted_at"`
}
