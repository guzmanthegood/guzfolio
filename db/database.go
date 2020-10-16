package db

import (
	"context"
	"log"

	"github.com/go-pg/pg/v10"
)

type logger struct{}

func (l logger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

func (l logger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	log.Println(q.FormattedQuery())
	return nil
}

func New(connectionString string, setLogger bool) (*pg.DB, error) {
	opt, err := pg.ParseURL(connectionString)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)

	if setLogger {
		db.AddQueryHook(logger{})
	}

	return db, nil
}