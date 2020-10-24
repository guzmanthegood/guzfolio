package server

import (
	"context"
	"log"
	"net/http"

	"guzfolio/datastore"
)

type server struct {
	server *http.Server
	port   string
	debug  bool
}

type Server interface {
	Start()
	Stop()
}

func CreateServer(debug bool, port string, ds datastore.DataStore) Server {
	return &server{
		server: &http.Server{
			Addr:    ":" + port,
			Handler: newRouter(ds),
		},
	}
}

// start server
func (s *server) Start() {
	log.Printf("[INFO] connect to http://localhost:%s/ for GraphQL playground", s.port)

	if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal("[ERRO] internal server error", err)
	}
}

// stop server
func (s *server) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Fatal("[ERRO] error during shutdown", err)
	}
}
