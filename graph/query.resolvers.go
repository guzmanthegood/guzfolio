package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"guzfolio/auth"
	"guzfolio/graph/generated"
	"guzfolio/model"
	"strconv"
)

func (r *queryResolver) Profile(ctx context.Context) (*model.User, error) {
	return r.DS.GetUserByID(auth.ContextAuthUser(ctx).UserID)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	userID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	a := auth.ContextAuthUser(ctx)
	if !auth.ContextAuthUser(ctx).IsAdmin && uint(userID) != a.UserID {
		return nil, errors.New("action not allowed")
	}
	return r.DS.GetUserByID(uint(userID))
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	if !auth.ContextAuthUser(ctx).IsAdmin {
		return nil, errors.New("action not allowed")
	}
	return r.DS.GetAllUsers()
}

func (r *queryResolver) AllCurrencies(ctx context.Context) ([]*model.Currency, error) {
	return r.DS.GetAllCurrencies()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
