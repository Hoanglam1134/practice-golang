package main

import (
	"database/sql"
	"log"
	"net"
	"practice-golang/inventory/api"
	inventory "practice-golang/inventory/sqlc"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:postgres@localhost:5432/inventory?sslmode=disable"
	serverAddress = "0.0.0.0:9090"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannnot connect to Db:", err)
	}

	store := inventory.NewStore(conn)
	server := api.NewServer(store)

	grpcServer := grpc.NewServer()
	api.RegisterInventoryServer(grpcServer, server)

	listener, err := net.Listen("tcp", serverAddress)

	if err != nil {
		log.Fatal("Cannot create listener")
	}

	log.Printf("Starting gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)

	if err != nil {
		log.Fatal("cannot serve gRPC server")
	}
}
