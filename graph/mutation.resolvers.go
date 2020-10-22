package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"guzfolio/auth"
	"guzfolio/graph/generated"
	"guzfolio/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	if !auth.ContextAuthUser(ctx).IsAdmin {
		return nil, errors.New("action not allowed")
	}
	return r.DS.CreateUser(input)
}

func (r *mutationResolver) CreatePortfolio(ctx context.Context, input model.CreatePortfolioInput) (*model.Portfolio, error) {
	if !auth.ContextAuthUser(ctx).IsAdmin {
		return nil, errors.New("action not allowed")
	}
	return r.DS.CreatePortfolio(input)
}

func (r *mutationResolver) CreateCurrency(ctx context.Context, input model.CreateCurrencyInput) (*model.Currency, error) {
	if !auth.ContextAuthUser(ctx).IsAdmin {
		return nil, errors.New("action not allowed")
	}
	return r.DS.CreateCurrency(input)
}

func (r *mutationResolver) CreateTransaction(ctx context.Context, input model.CreateTransactionInput) (*model.Transaction, error) {
	if !auth.ContextAuthUser(ctx).IsAdmin {
		return nil, errors.New("action not allowed")
	}
	return r.DS.CreateTransaction(input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
