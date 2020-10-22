//go:generate go run github.com/vektah/dataloaden UserLoader uint *guzfolio/model.User
//go:generate go run github.com/vektah/dataloaden CurrencyLoader uint *guzfolio/model.Currency
//go:generate go run github.com/vektah/dataloaden PortfoliosLoader uint []*guzfolio/model.Portfolio
//go:generate go run github.com/vektah/dataloaden TransactionsLoader uint []*guzfolio/model.Transaction

package dataloader

import (
	"context"
	"net/http"
	"time"

	"guzfolio/datastore"
	"guzfolio/model"
)

type Loaders struct {
	UserByID					*UserLoader
	CurrencyByID				*CurrencyLoader
	PortfoliosByUser			*PortfoliosLoader
	TransactionsByPortfolio		*TransactionsLoader
}

func Middleware(ds datastore.DataStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := Loaders{}

		// set this to zero what happens without dataloading
		wait := 250 * time.Microsecond

		// simple 1:1 loader, fetch an user by its primary key
		loaders.UserByID = &UserLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []uint) ([]*model.User, []error) {
				users, err := ds.GetUserByIDs(keys)
				if err != nil {
					return nil, []error{err}
				}
				return users, nil
			},
		}

		// simple 1:1 loader, fetch an currency by its primary key
		loaders.CurrencyByID = &CurrencyLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []uint) ([]*model.Currency, []error) {
				currencies, err := ds.GetCurrencyByIDs(keys)
				if err != nil {
					return nil, []error{err}
				}
				return currencies, nil
			},
		}

		// 1:M loader, fetch portfolios by user key
		loaders.PortfoliosByUser = &PortfoliosLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []uint) ([][]*model.Portfolio, []error) {
				portfoliosDB, err := ds.GetPortfoliosByUserIDs(keys)
				if err != nil {
					return nil, []error{err}
				}

				portfoliosMap := make(map[uint][]*model.Portfolio)
				for _, p := range portfoliosDB {
					portfoliosMap[p.UserID] = append(portfoliosMap[p.UserID], p)
				}

				portfolios := make([][]*model.Portfolio, len(keys))
				for i, k := range keys {
					if p, ok := portfoliosMap[k]; ok {
						portfolios[i] = p
					}
				}
				return portfolios, nil
			},
		}

		// 1:M loader, fetch transactions by profile key
		loaders.TransactionsByPortfolio = &TransactionsLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []uint) ([][]*model.Transaction, []error) {
				transactionsDB, err := ds.GetTransactionsByPortfolioIDs(keys)
				if err != nil {
					return nil, []error{err}
				}

				transactionsMap := make(map[uint][]*model.Transaction)
				for _, t := range transactionsDB {
					transactionsMap[t.PortfolioID] = append(transactionsMap[t.PortfolioID], t)
				}

				transactions := make([][]*model.Transaction, len(keys))
				for i, k := range keys {
					if p, ok := transactionsMap[k]; ok {
						transactions[i] = p
					}
				}
				return transactions, nil
			},
		}

		dlCtx := context.WithValue(r.Context(), contextKey, loaders)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

type contextKeyType struct{ name string }

var contextKey = contextKeyType{"userContext"}

func ContextLoaders(ctx context.Context) Loaders {
	return ctx.Value(contextKey).(Loaders)
}