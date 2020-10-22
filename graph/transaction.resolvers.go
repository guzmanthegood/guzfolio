package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"guzfolio/datastore/dataloader"
	"guzfolio/graph/generated"
	"guzfolio/model"
	"strconv"
)

func (r *transactionResolver) ID(ctx context.Context, obj *model.Transaction) (string, error) {
	return strconv.FormatUint(uint64(obj.ID), 10), nil
}

func (r *transactionResolver) Currency(ctx context.Context, obj *model.Transaction) (*model.Currency, error) {
	return dataloader.ContextLoaders(ctx).CurrencyByID.Load(obj.CurrencyID)
}

func (r *transactionResolver) Portfolio(ctx context.Context, obj *model.Transaction) (*model.Portfolio, error) {
	return dataloader.ContextLoaders(ctx).PortfolioByID.Load(obj.PortfolioID)
}

// Transaction returns generated.TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

type transactionResolver struct{ *Resolver }
