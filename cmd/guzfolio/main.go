package main

import (
	"os"

	"guzfolio/datastore/postgres"
	"guzfolio/server"
)

const defaultPort = "8080"

func main() {

	// get default port
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = defaultPort
	}

	// get debug mode
	serverDebug := true
	if os.Getenv("DEBUG") == "0" {
		serverDebug = false
	}

	// get default postgres connection
	pgConnectionString := os.Getenv("PG_CONNECTION_STRING")
	if pgConnectionString == "" {
		pgConnectionString = "postgres://user:pass@localhost:5432/db_name?sslmode=disable"
	}

	ds := postgres.New(pgConnectionString, serverDebug)
	s := server.CreateServer(serverDebug, serverPort, ds)
	s.Start()
}
