package postgres

import "guzfolio/model"

func (ds dataStore) GetUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	result := ds.db.First(user, id)
	return user, result.Error
}

func (ds dataStore) GetUserByIDs(ids []uint) ([]*model.User, error) {
	var users []*model.User
	result := ds.db.Find(&users, ids)
	return users, result.Error
}

func (ds dataStore) GetUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	result := ds.db.Where("email = ?", email).First(&user)
	return user, result.Error
}

func (ds dataStore) GetCurrencyByID (id uint) (*model.Currency, error) {
	currency := &model.Currency{}
	result := ds.db.First(currency, id)
	return currency, result.Error
}

func (ds dataStore) GetCurrencyByIDs (ids []uint) ([]*model.Currency, error) {
	var currencies []*model.Currency
	result := ds.db.Find(&currencies, ids)
	return currencies, result.Error
}

func (ds dataStore) GetTransactionsByPortfolioID (id uint) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	result := ds.db.Where("portfolio_id = ?", id).Find(&transactions)
	return transactions, result.Error
}

func (ds dataStore) GetTransactionsByPortfolioIDs (ids []uint) ([]*model.Transaction, error) {
	var transactions []*model.Transaction
	result := ds.db.Where("portfolio_id IN ?", ids).Find(&transactions)
	return transactions, result.Error
}

func (ds dataStore) GetAllCurrencies() ([]*model.Currency, error) {
	var currencies []*model.Currency
	result := ds.db.Find(&currencies)
	return currencies, result.Error
}

func (ds dataStore) GetAllUsers() ([]*model.User, error) {
	var users []*model.User
	result := ds.db.Find(&users)
	return users, result.Error
}

func (ds dataStore) GetPortfolioByID(id uint) (*model.Portfolio, error) {
	portfolio := &model.Portfolio{}
	result := ds.db.First(portfolio, id)
	return portfolio, result.Error
}

func (ds dataStore) GetPortfolioByIDs(ids []uint) ([]*model.Portfolio, error) {
	var portfolio []*model.Portfolio
	result := ds.db.Find(&portfolio, ids)
	return portfolio, result.Error
}

func (ds dataStore) GetPortfoliosByUserID(id uint) ([]*model.Portfolio, error) {
	var portfolios []*model.Portfolio
	result := ds.db.Where("user_id = ?", id).Find(&portfolios)
	return portfolios, result.Error
}

func (ds dataStore) GetPortfoliosByUserIDs(ids []uint) ([]*model.Portfolio, error) {
	var portfolios []*model.Portfolio
	result := ds.db.Where("user_id IN ?", ids).Find(&portfolios)
	return portfolios, result.Error
}
