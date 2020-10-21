package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"guzfolio/auth"
	"guzfolio/graph/generated"
	"guzfolio/model"
	"strconv"
)

func (r *queryResolver) Profile(ctx context.Context) (*model.User, error) {
	u := auth.ContextAuthUser(ctx)
	return r.DS.GetUserByID(u.UserID)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	return r.DS.GetUserByID(uint(userID))
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	return r.DS.GetAllUsers()
}

func (r *queryResolver) AllCurrencies(ctx context.Context) ([]*model.Currency, error) {
	return r.DS.GetAllCurrencies()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
