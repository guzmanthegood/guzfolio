package postgres

import (
	"guzfolio/datastore"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type dataStore struct {
	db *gorm.DB
}

func New(dsn string, debug bool) datastore.DataStore  {
	loggerLvl := logger.Silent
	if debug {
		loggerLvl = logger.Info
	}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(loggerLvl),
	})

	if err != nil {
		panic(err)
	}
	return dataStore{db: db}
}
