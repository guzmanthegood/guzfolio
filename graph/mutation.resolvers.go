package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"guzfolio/graph/generated"
	"guzfolio/model"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	//// secure password
	//bytePassword := []byte(input.Password)
	//passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	//if err != nil {
	//	return nil, err
	//}
	//input.Password = string(passwordHash)
	//
	//// create user
	//user, err := r.DS.RegisterUser(input)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// generate new access token
	//token, err := auth.GenerateToken(user.ID, user.Email, false)
	//if err != nil {
	//	return nil, err
	//}
	//
	//authResponse := &model.AuthResponse{
	//	AuthToken: &model.AuthToken{
	//		AccessToken: token.AccessToken,
	//		ExpiredAt:   token.ExpiredAt,
	//	},
	//	User: user,
	//}
	//return authResponse, nil
	return nil, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	//// get user
	//user, err := r.DS.GetUserByEmail(input.Email)
	//if err != nil {
	//	return nil, err
	//}
	//
	//// compare password
	//bytePassword := []byte(input.Password)
	//byteHashedPassword := []byte(user.Password)
	//if err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword); err != nil {
	//	return nil, err
	//}
	//
	//// generate new access token
	//token, err := auth.GenerateToken(user.ID, user.Email, false)
	//if err != nil {
	//	return nil, err
	//}
	//
	//authResponse := &model.AuthResponse{
	//	AuthToken: &model.AuthToken{
	//		AccessToken: token.AccessToken,
	//		ExpiredAt:   token.ExpiredAt,
	//	},
	//	User: user,
	//}
	//return authResponse, nil
	return nil, nil
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
