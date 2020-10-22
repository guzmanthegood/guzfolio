package datastore

import (
	"guzfolio/model"
)

type DataStore interface {
	GetUserByID(id uint) (*model.User, error)
	GetUserByIDs(ids []uint) ([]*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetCurrencyByID (id uint) (*model.Currency, error)
	GetCurrencyByIDs (ids []uint) ([]*model.Currency, error)
	GetTransactionsByPortfolioID (id uint) ([]*model.Transaction, error)
	GetTransactionsByPortfolioIDs (ids []uint) ([]*model.Transaction, error)
	GetAllCurrencies() ([]*model.Currency, error)
	GetAllUsers() ([]*model.User, error)
	GetPortfolioByID(id uint) (*model.Portfolio, error)
	GetPortfolioByIDs(ids []uint) ([]*model.Portfolio, error)
	GetPortfoliosByUserID(id uint) ([]*model.Portfolio, error)
	GetPortfoliosByUserIDs(id []uint) ([]*model.Portfolio, error)
	GetPortfolioReport(id uint) (model.PortfolioReport, error)
	CreateUser(input model.CreateUserInput) (*model.User, error)
	CreatePortfolio(input model.CreatePortfolioInput) (*model.Portfolio, error)
	CreateCurrency(input model.CreateCurrencyInput) (*model.Currency, error)
	CreateTransaction(input model.CreateTransactionInput) (*model.Transaction, error)
}