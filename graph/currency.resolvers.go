package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"guzfolio/graph/generated"
	"guzfolio/model"
	"strconv"
)

func (r *currencyResolver) ID(ctx context.Context, obj *model.Currency) (string, error) {
	return strconv.FormatUint(uint64(obj.ID), 10), nil
}

// Currency returns generated.CurrencyResolver implementation.
func (r *Resolver) Currency() generated.CurrencyResolver { return &currencyResolver{r} }

type currencyResolver struct{ *Resolver }
