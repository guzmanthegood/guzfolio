package model

type Portfolio struct {
	ID           string    `json:"id"`
	Name         *string   `json:"name"`
	FiatCurrency *Currency `json:"fiatCurrency"`
	User         *User     `json:"user"`
}
