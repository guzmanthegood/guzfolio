package main

import (
	"guzfolio/datastore"
	"guzfolio/graph"
	"guzfolio/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
)

func newRouter(ds datastore.DataStore) *chi.Mux {
	r := chi.NewRouter()

	schema := generated.NewExecutableSchema(generated.Config{
		Resolvers: &graph.Resolver{
			DS: ds,
		}},
	)
	srv := handler.NewDefaultServer(schema)

	r.Handle("/", playground.Handler("GraphQL playground", "/query"))
	r.Handle("/query", srv)

	return r
}
