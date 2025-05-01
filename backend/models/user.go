package models

type User struct {
	// gorm.Model  `json:"-"`
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `gorm:"unique"`
	Password []byte `json:"-"`
}
