package model

type Portfolio struct {
	ID           	uint      	`json:"id"`
	Name         	*string   	`json:"name"`
	User         	User     	`json:"user"`
	UserID         	uint	  	`json:"userID"`
}