package api

import (
	inventory "practice-golang/inventory/sqlc"
)

type Server struct {
	UnimplementedInventoryServer
	store *inventory.Store
}

func NewServer(store *inventory.Store) *Server {
	server := &Server{store: store}
	return server
}
