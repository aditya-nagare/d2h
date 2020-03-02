package models

//Service Model
type Service struct {
	ID    int    `gorm:"id"`
	Name  string `gorm:"name"`
	Price uint   `gorm:"price"`
}
