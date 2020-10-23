package datastore

import (
	"guzfolio/model"
)

type DataStore interface {
	GetUserByIDs(ids []uint) ([]*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetCurrencyByIDs (ids []uint) ([]*model.Currency, error)
	GetTransactionsByPortfolioIDs (ids []uint) ([]*model.Transaction, error)
	GetAllCurrencies() ([]*model.Currency, error)
	GetAllUsers() ([]*model.User, error)
	GetPortfolioByIDs(ids []uint) ([]*model.Portfolio, error)
	GetPortfoliosByUserIDs(id []uint) ([]*model.Portfolio, error)
	GetPortfolioReports(ids []uint) ([]*model.PortfolioReport, error)
	CreateUser(input model.CreateUserInput) (*model.User, error)
	CreatePortfolio(input model.CreatePortfolioInput) (*model.Portfolio, error)
	CreateCurrency(input model.CreateCurrencyInput) (*model.Currency, error)
	CreateTransaction(input model.CreateTransactionInput) (*model.Transaction, error)
}