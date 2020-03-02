package models

//Channel Model
type Channel struct {
	ID    int    `gorm:"id"`
	Name  string `gorm:"name"`
	Price uint   `gorm:"price"`
}
