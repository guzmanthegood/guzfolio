//go:generate go run github.com/vektah/dataloaden CurrencyLoader uint *guzfolio/model.Currency

package dataloader

import (
	"context"
	"net/http"
	"time"

	"guzfolio/datastore"
	"guzfolio/model"
)

type Loaders struct {
	CurrencyByID	*CurrencyLoader
}

func LoaderMiddleware(ds datastore.DataStore, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loaders := Loaders{}

		// set this to zero what happens without dataloading
		wait := 250 * time.Microsecond

		// simple 1:1 loader, fetch an address by its primary key
		loaders.CurrencyByID = &CurrencyLoader{
			wait:     wait,
			maxBatch: 100,
			fetch: func(keys []uint) ([]*model.Currency, []error) {
				errors := make([]error, len(keys))
				currencies, err := ds.GetCurrencyByIDs(keys)
				if err != nil {
					errors = append(errors, err)
				}
				return currencies, errors
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