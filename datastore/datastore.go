package datastore

import (
	"guzfolio/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DataStore interface {
	AutoMigrate() error
	GetDB() *gorm.DB
	GetUserByID(id string) (*model.User, error)
}

type dataStore struct {
	db *gorm.DB
}

func New(dsn string) DataStore  {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return dataStore{db: db}
}

func (ds dataStore) GetUserByID(id string) (*model.User, error) {
	return nil, nil
}

func (ds dataStore) AutoMigrate() error {
	ds.db.AutoMigrate(&model.Currency{})
	ds.db.AutoMigrate(&model.User{})
	ds.db.AutoMigrate(&model.Portfolio{})
	return nil
}

func (ds dataStore) GetDB() *gorm.DB {
	return ds.db
}