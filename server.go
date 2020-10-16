package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"guzfolio/db"
	"guzfolio/graph"
	"guzfolio/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-pg/pgext"
)

const defaultPort = "8080"

func main() {
	// create new db connection
	DB, err := db.New(os.Getenv("DB_CONNECTION_STRING"), false)
	defer DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	DB.AddQueryHook(&pgext.OpenTelemetryHook{})

	// ping db connection & create schema
	ctx := context.Background()
	if err := DB.Ping(ctx); err != nil {
		log.Fatal(err)
	}
	if err := db.CreateSchema(DB); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
