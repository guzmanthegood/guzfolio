package model

type User struct {
	ID         uint      	`json:"id"`
	Email      string       `json:"email" gorm:"unique"`
	Password   string       `json:"password"`
	Name       string       `json:"name"`
	Portfolios []*Portfolio `json:"portfolios"`
}