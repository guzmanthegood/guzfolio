package datastore

import (
	"guzfolio/model"
)

type DataStore interface {
	GetUserByID(id uint) (*model.User, error)
	GetCurrencyByID (id uint) (*model.Currency, error)
	GetAllCurrencies() ([]*model.Currency, error)
	GetAllUsers() ([]*model.User, error)
	GetPortfoliosByUserID(id uint) ([]*model.Portfolio, error)
	CreateUser(input model.CreateUserInput) (*model.User, error)
	CreatePortfolio(input model.CreatePortfolioInput) (*model.Portfolio, error)
	CreateCurrency(input model.CreateCurrencyInput) (*model.Currency, error)
}