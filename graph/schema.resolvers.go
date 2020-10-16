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
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateCurrency(ctx context.Context, input model.CreateCurrencyInput) (*model.Currency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *portfolioResolver) FiatCurrency(ctx context.Context, obj *model.Portfolio) (*model.Currency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *portfolioResolver) User(ctx context.Context, obj *model.Portfolio) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Portfolio(ctx context.Context, id string) (*model.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllCurrencies(ctx context.Context) ([]*model.Currency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Portfolios(ctx context.Context, obj *model.User) ([]*model.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Portfolio returns generated.PortfolioResolver implementation.
func (r *Resolver) Portfolio() generated.PortfolioResolver { return &portfolioResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type portfolioResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
