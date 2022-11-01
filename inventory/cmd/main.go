package main

import (
	"database/sql"
	"log"
	"todolist-grpc/inventory/api"
	inventory "todolist-grpc/inventory/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:postgres@localhost:5432/inventory?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannnot connect to Db:", err)
	}

	store := inventory.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Cannot start the server:", err)
	}
}
