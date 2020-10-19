//go:generate go run github.com/vektah/dataloaden CurrencyLoader uint *guzfolio/model.Currency
//go:generate go run github.com/vektah/dataloaden PortfoliosLoader uint []*guzfolio/model.Portfolio

package dataloader

import (
	"context"
	"net/http"
	"time"

	"guzfolio/datastore"
	"guzfolio/model"
)

type Loaders struct {
	CurrencyByID		*CurrencyLoader
	PortfoliosByUser	*PortfoliosLoader
}

func LoaderMiddleware(ds datastore.DataStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := Loaders{}

		// set this to zero what happens without dataloading
		wait := 250 * time.Microsecond

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

		dlCtx := context.WithValue(r.Context(), contextKey, loaders)
		next.ServeHTTP(w, r.WithContext(dlCtx))
	})
}

type contextKeyType struct{ name string }

var contextKey = contextKeyType{"userContext"}

func ContextLoaders(ctx context.Context) Loaders {
	return ctx.Value(contextKey).(Loaders)
}