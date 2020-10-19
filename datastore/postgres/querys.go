package postgres

import "guzfolio/model"

func (ds dataStore) GetUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	result := ds.db.First(user, id)
	return user, result.Error
}
