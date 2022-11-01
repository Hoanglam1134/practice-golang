package api

import (
	inventory "todolist-grpc/inventory/sqlc"
)

type Server struct {
	UnimplementedInventoryServer
	store *inventory.Store
}

func NewServer(store *inventory.Store) *Server {
	server := &Server{store: store}
	return server
}
