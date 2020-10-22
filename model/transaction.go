package model

import "time"

type Transaction struct {
	ID           uint     	`json:"id"`
	PricePerCoin float64    `json:"pricePerCoin"`
	Quantity     float64    `json:"quantity"`
	Currency     Currency   `json:"currency"`
	CurrencyID	 uint		`json:"currencyID"`
	Portfolio    Portfolio  `json:"portfolio"`
	PortfolioID	 uint		`json:"portfolioID"`
	UpdatedAt    time.Time 	`json:"updatedAt"`
	CreatedAt    time.Time 	`json:"createdAt"`
}
