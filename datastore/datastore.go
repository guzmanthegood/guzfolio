package datastore

import (
	"guzfolio/model"
)

type DataStore interface {
	GetUserByID(id uint) (*model.User, error)
	CreateUser(input model.CreateUserInput) (*model.User, error)
	CreatePortfolio(input model.CreatePortfolioInput) (*model.Portfolio, error)
	CreateCurrency(input model.CreateCurrencyInput) (*model.Currency, error)
}