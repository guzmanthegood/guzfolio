package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"guzfolio/graph/generated"
	"guzfolio/model"
	"strconv"
)

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return strconv.FormatUint(uint64(obj.ID), 10), nil
}

func (r *userResolver) Portfolios(ctx context.Context, obj *model.User) ([]*model.Portfolio, error) {
	return r.DS.GetPortfoliosByUserID(obj.ID)
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
