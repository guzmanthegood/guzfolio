package db

import (
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"guzfolio/model"
)

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*model.Currency)(nil),
		(*model.User)(nil),
		(*model.Portfolio)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists:   true,
			Temp:          false,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}