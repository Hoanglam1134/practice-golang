package main

import (
	"database/sql"
	"log"
	"net"
	"practice-golang/service-order/adapter"
	"practice-golang/service-order/api"
	"practice-golang/service-order/server"
	orders "practice-golang/service-order/sqlc"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgres://postgres:postgres@localhost:5432/orders?sslmode=disable"
	serverAddress = "0.0.0.0:10987"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannnot connect to Db:", err)
	}

	cc := adapter.NewClientInventory()
	store := orders.NewStore(conn)
	server := server.NewServer(store, cc)

	grpcServer := grpc.NewServer()
	api.RegisterOrdersServer(grpcServer, server)

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
