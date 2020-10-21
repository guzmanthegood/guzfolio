package postgres

import "guzfolio/model"

func (ds dataStore) GetUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	result := ds.db.First(user, id)
	return user, result.Error
}

func (ds dataStore) GetUserByIDs(ids []uint) ([]*model.User, error) {
	users := []*model.User{}
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
	currencies := []*model.Currency{}
	result := ds.db.Find(&currencies, ids)
	return currencies, result.Error
}

func (ds dataStore) GetTransactionsByUserID (id uint) ([]*model.Transaction, error){
	transactions := []*model.Transaction{}
	result := ds.db.Where("portfolio_id = ?", id).Find(&transactions)
	return transactions, result.Error
}

func (ds dataStore) GetAllCurrencies() ([]*model.Currency, error) {
	currencies := []*model.Currency{}
	result := ds.db.Find(&currencies)
	return currencies, result.Error
}

func (ds dataStore) GetAllUsers() ([]*model.User, error){
	users := []*model.User{}
	result := ds.db.Find(&users)
	return users, result.Error
}

func (ds dataStore) GetPortfoliosByUserID(id uint) ([]*model.Portfolio, error){
	portfolios := []*model.Portfolio{}
	result := ds.db.Where("user_id = ?", id).Find(&portfolios)
	return portfolios, result.Error
}

func (ds dataStore) GetPortfoliosByUserIDs(ids []uint) ([]*model.Portfolio, error){
	portfolios := []*model.Portfolio{}
	result := ds.db.Where("user_id IN ?", ids).Find(&portfolios)
	return portfolios, result.Error
}
