package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"guzfolio/datastore/dataloader"
	"guzfolio/graph/generated"
	"guzfolio/model"
	"strconv"
)

func (r *portfolioResolver) ID(ctx context.Context, obj *model.Portfolio) (string, error) {
	return strconv.FormatUint(uint64(obj.ID), 10), nil
}

func (r *portfolioResolver) TotalValue(ctx context.Context, obj *model.Portfolio) (float64, error) {
	report, err := dataloader.ContextLoaders(ctx).PortfolioReportByID.Load(obj.ID)
	if report == nil {
		return 0, err
	}
	return report.TotalValue, err
}

func (r *portfolioResolver) NetCost(ctx context.Context, obj *model.Portfolio) (float64, error) {
	report, err := dataloader.ContextLoaders(ctx).PortfolioReportByID.Load(obj.ID)
	if report == nil {
		return 0, err
	}
	return report.NetCost, err
}

func (r *portfolioResolver) PercentChange(ctx context.Context, obj *model.Portfolio) (float64, error) {
	report, err := dataloader.ContextLoaders(ctx).PortfolioReportByID.Load(obj.ID)
	if report == nil {
		return 0, err
	}
	return report.PercentChange, err
}

func (r *portfolioResolver) TotalTransactions(ctx context.Context, obj *model.Portfolio) (int, error) {
	report, err := dataloader.ContextLoaders(ctx).PortfolioReportByID.Load(obj.ID)
	if report == nil {
		return 0, err
	}
	return report.TotalTransactions, err
}

func (r *portfolioResolver) User(ctx context.Context, obj *model.Portfolio) (*model.User, error) {
	return dataloader.ContextLoaders(ctx).UserByID.Load(obj.UserID)
}

func (r *portfolioResolver) Transactions(ctx context.Context, obj *model.Portfolio) ([]*model.Transaction, error) {
	return dataloader.ContextLoaders(ctx).TransactionsByPortfolio.Load(obj.ID)
}

// Portfolio returns generated.PortfolioResolver implementation.
func (r *Resolver) Portfolio() generated.PortfolioResolver { return &portfolioResolver{r} }

type portfolioResolver struct{ *Resolver }
