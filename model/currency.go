package model

type Currency struct {
	ID   uint 	`json:"id"`
	Code string `json:"code" gorm:"unique"`
	Name string `json:"name"`
}