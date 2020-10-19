package model

type Portfolio struct {
	ID           	uint      	`json:"id"`
	Name         	*string   	`json:"name"`
	FiatCurrency 	Currency 	`json:"fiatCurrency"`
	FiatCurrencyID 	uint	  	`json:"fiatCurrencyID"`
	User         	User     	`json:"user"`
	UserID         	uint	  	`json:"userID"`
}