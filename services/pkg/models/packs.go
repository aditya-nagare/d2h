package models

//Pack Model
type Pack struct {
	ID    int    `gorm:"id"`
	Name  string `gorm:"name"`
	Price uint   `gorm:"price"`
}

//PackComposition Model
type PackComposition struct {
	ID   int    `gorm:"id"`
	Name string `gorm:"name"`
}
