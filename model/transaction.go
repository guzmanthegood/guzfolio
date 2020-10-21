package model

import "time"

type Transaction struct {
	ID           uint     	`json:"id"`
	BoughtWith   *Currency  `json:"boughtWith"`
	BoughtWithID uint		`json:"boughtWithID"`
	PricePerCoin float64    `json:"pricePerCoin"`
	Quantity     float64    `json:"quantity"`
	Currency     *Currency  `json:"currency"`
	CurrencyID	 uint		`json:"currencyID"`
	Date         time.Time  `json:"date"`
	Portfolio    *Portfolio `json:"portfolio"`
}
