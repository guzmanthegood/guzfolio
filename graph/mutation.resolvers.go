package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"guzfolio/graph/generated"
	"guzfolio/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	user := &model.User{
		Email: input.Email,
		Name:  input.Name,
	}
	result := r.DS.GetDB().Create(user)
	return user, result.Error
}

func (r *mutationResolver) CreatePortfolio(ctx context.Context, input *model.CreatePortfolioInput) (*model.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCurrency(ctx context.Context, input model.CreateCurrencyInput) (*model.Currency, error) {
	currency := &model.Currency{
		Code: input.Code,
		Name: input.Name,
		Type: input.Type,
	}
	result := r.DS.GetDB().Create(currency)
	return currency, result.Error
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
