package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"guzfolio/auth"
	"guzfolio/graph/generated"
	"guzfolio/model"

	"golang.org/x/crypto/bcrypt"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	return r.DS.CreateUser(input)
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	// secure password
	bytePassword := []byte(input.Password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	input.Password = string(passwordHash)

	// create user
	user, err := r.DS.RegisterUser(input)
	if err != nil {
		return nil, err
	}

	// generate new access token
	token, err := auth.GenerateToken(user.ID, user.Email, false)
	if err != nil {
		return nil, err
	}

	authResponse := &model.AuthResponse{
		AuthToken: &model.AuthToken{
			AccessToken: token.AccessToken,
			ExpiredAt:   token.ExpiredAt,
		},
		User: user,
	}
	return authResponse, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePortfolio(ctx context.Context, input model.CreatePortfolioInput) (*model.Portfolio, error) {
	return r.DS.CreatePortfolio(input)
}

func (r *mutationResolver) CreateCurrency(ctx context.Context, input model.CreateCurrencyInput) (*model.Currency, error) {
	return r.DS.CreateCurrency(input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
