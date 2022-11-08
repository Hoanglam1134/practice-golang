package server

import (
	"log"
	"practice-golang/inventory/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClientInventory() (api.InventoryClient, *grpc.ClientConn) {
	cc, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	client := api.NewInventoryClient(cc)
	return client, cc
}
