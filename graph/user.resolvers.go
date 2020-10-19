package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"guzfolio/graph/generated"
	"guzfolio/model"
)

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Portfolios(ctx context.Context, obj *model.User) ([]*model.Portfolio, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }