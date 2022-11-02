package server

import (
	"log"
	"practice-golang/inventory/api"

	"google.golang.org/grpc"
)

func NewClientInventory() api.InventoryClient {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	client := api.NewInventoryClient(cc)
	return client
}
