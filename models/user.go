package models

type User struct {
	ID       uint   `json:"id" gorm:"primary-key"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}
