package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"guzfolio/datastore/dataloader"
	"guzfolio/graph/generated"
	"guzfolio/model"
)

func (r *authResponseResolver) User(ctx context.Context, obj *model.AuthResponse) (*model.User, error) {
	return dataloader.ContextLoaders(ctx).UserByID.Load(obj.User.ID)
}

// AuthResponse returns generated.AuthResponseResolver implementation.
func (r *Resolver) AuthResponse() generated.AuthResponseResolver { return &authResponseResolver{r} }

type authResponseResolver struct{ *Resolver }

