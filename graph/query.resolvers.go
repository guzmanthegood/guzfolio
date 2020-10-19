package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"guzfolio/graph/generated"
	"guzfolio/model"
	"log"
)

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	log.Println("query user ************")
	user := &model.User{}
	r.DS.GetDB().First(user, 10)
	return user, nil
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
