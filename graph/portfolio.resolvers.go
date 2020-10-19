package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"guzfolio/graph/generated"
	"guzfolio/model"
)

func (r *portfolioResolver) ID(ctx context.Context, obj *model.Portfolio) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *portfolioResolver) FiatCurrency(ctx context.Context, obj *model.Portfolio) (*model.Currency, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *portfolioResolver) User(ctx context.Context, obj *model.Portfolio) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Portfolio returns generated.PortfolioResolver implementation.
func (r *Resolver) Portfolio() generated.PortfolioResolver { return &portfolioResolver{r} }

type portfolioResolver struct{ *Resolver }
