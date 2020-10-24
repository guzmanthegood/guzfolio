package server

import (
	"guzfolio/auth"
	"guzfolio/datastore"
	"guzfolio/datastore/dataloader"
	"guzfolio/graph"
	"guzfolio/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func newRouter(ds datastore.DataStore) *chi.Mux {
	r := chi.NewRouter()

	queryHandler := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			DS: ds,
		}},
	))

	r.Mount("/auth", auth.Router(ds))
	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", auth.Middleware(dataloader.Middleware(ds, queryHandler)))

	return r
}
