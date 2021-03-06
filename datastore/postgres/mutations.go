package postgres

import (
	"golang.org/x/crypto/bcrypt"
	"guzfolio/model"
)

func (ds dataStore) CreateUser(input model.CreateUserInput) (*model.User, error) {
	// secure password
	bytePassword := []byte(input.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &model.User{
		Email: input.Email,
		Password: string(passwordHash),
		Name:  input.Name,
	}
	result := ds.db.Create(user)
	return user, result.Error
}

func (ds dataStore) CreatePortfolio(input model.CreatePortfolioInput) (*model.Portfolio, error) {
	// get user
	user := model.User{}
	result := ds.db.Where("email = ?", input.UserEmail).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}


	// create portfolio
	portfolio := &model.Portfolio{
		User:       	user,
		Name:           input.Name,
	}
	result = ds.db.Create(portfolio)
	return portfolio, result.Error
}

func (ds dataStore) CreateCurrency(input model.CreateCurrencyInput) (*model.Currency, error) {
	currency := &model.Currency{
		Code: input.Code,
		Name: input.Name,
		MarketValue: input.MarketValue,
	}
	result := ds.db.Create(currency)
	return currency, result.Error
}

func (ds dataStore) CreateTransaction(input model.CreateTransactionInput) (*model.Transaction, error) {
	// get portfolio
	portfolio := model.Portfolio{}
	result := ds.db.First(&portfolio, input.PorfolioID)
	if result.Error != nil {
		return nil, result.Error
	}

	// get currency
	currency := model.Currency{}
	result = ds.db.Where("code = ?", input.CurrencyCode).First(&currency)
	if result.Error != nil {
		return nil, result.Error
	}

	// create transaction
	transaction := &model.Transaction{
		PricePerCoin: input.PricePerCoin,
		Quantity:     input.Quantity,
		Currency:     currency,
		Portfolio:    portfolio,
	}
	result = ds.db.Create(transaction)

	return transaction, result.Error
}