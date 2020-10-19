package model

type User struct {
	ID         uint      	`json:"id"`
	Email      string       `json:"email" gorm:"unique"`
	Name       string       `json:"name"`
	Portfolios []*Portfolio `json:"portfolios"`
}
