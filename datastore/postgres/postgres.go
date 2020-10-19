package postgres

import (
	"guzfolio/datastore"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataStore struct {
	db *gorm.DB
}

func New(dsn string) datastore.DataStore  {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return dataStore{db: db}
}
