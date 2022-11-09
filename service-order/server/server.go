package server

import (
	"practice-golang/service-order/adapter"
	"practice-golang/service-order/api"
	orders "practice-golang/service-order/sqlc"
)

type Server struct {
	api.UnimplementedOrdersServer
	store           orders.StoreQuerier
	inventoryClient adapter.ClientAdapter
}

func NewServer(store orders.StoreQuerier, client adapter.ClientAdapter) *Server {
	server := &Server{
		store:           store,
		inventoryClient: client,
	}
	return server
}
