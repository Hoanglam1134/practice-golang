package server

import (
	"practice-golang/service-order/api"
	orders "practice-golang/service-order/sqlc"
)

type Server struct {
	api.UnimplementedOrdersServer
	store *orders.Store
}

func NewServer(store *orders.Store) *Server {
	server := &Server{store: store}
	return server
}
