package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"guzfolio/datastore"
)

const defaultPort = "8080"

type server struct {
	server 	*http.Server
	port   	string
}

func main() {
	s := server{}
	s.initialize()
	s.start()
}

func (s *server) initialize() {
	// get default port
	s.port = os.Getenv("PORT")
	if s.port == "" {
		s.port = defaultPort
	}

	// new data base connection
	ds := datastore.New(os.Getenv("PG_CONNECTION_STRING"))
	err := ds.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}

	// initialize server
	s.server = &http.Server{
		Addr:    ":" + s.port,
		Handler: newRouter(ds),
	}
}

// start server
func (s *server) start() {
	log.Printf("[INFO] connect to http://localhost:%s/ for GraphQL playground", s.port)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("[ERRO] internal server error", err)
	}
}

// stop server
func (s *server) stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Fatal("[ERRO] error during shutdown", err)
	}
}
