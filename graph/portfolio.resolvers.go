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

func (r *portfolioResolver) ID(ctx context.Context, obj *model.Portfolio) (string, error) {
	return strconv.FormatUint(uint64(obj.ID), 10), nil
}

func (r *portfolioResolver) FiatCurrency(ctx context.Context, obj *model.Portfolio) (*model.Currency, error) {
	return r.DS.GetCurrencyByID(obj.FiatCurrencyID)
}

func (r *portfolioResolver) FiatCurrencyLoader(ctx context.Context, obj *model.Portfolio) (*model.Currency, error) {
	return dataloader.ContextLoaders(ctx).CurrencyByID.Load(obj.FiatCurrencyID)
}

func (r *portfolioResolver) User(ctx context.Context, obj *model.Portfolio) (*model.User, error) {
	return r.DS.GetUserByID(obj.UserID)
}

func (r *portfolioResolver) UserLoader(ctx context.Context, obj *model.Portfolio) (*model.User, error) {
	return dataloader.ContextLoaders(ctx).UserByID.Load(obj.UserID)
}

// Portfolio returns generated.PortfolioResolver implementation.
func (r *Resolver) Portfolio() generated.PortfolioResolver { return &portfolioResolver{r} }

type portfolioResolver struct{ *Resolver }
