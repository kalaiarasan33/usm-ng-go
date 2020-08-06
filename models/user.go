package models

type User struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Name   string `json:"name"`
	Depart string `json:"depart"`
}
