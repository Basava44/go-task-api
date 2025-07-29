package models

type Task struct {
	ID     uint   `json:"id" gorm:"primary-key"`
	Title  string `json:"title"`
	Status string `json:"status"`
	UserID uint   `json:"userId"`
}
