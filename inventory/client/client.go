package main

import (
	"context"
	"log"

	"practice-golang/inventory/api"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:9090", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	defer cc.Close()

	client := api.NewInventoryClient(cc)
	callGetQuantity(client)
	//callUpdateQuantity(client)
	log.Printf("service client %v", client)
}

func callGetQuantity(c api.InventoryClient) {
	log.Printf("calling service get")
	sku := "Tablet"
	resp, err := c.GetQuantity(context.Background(), &api.GetQuantityRequest{
		Sku: sku,
	})

	if err != nil {
		log.Fatalf("Error when calc sum: %v", err)
	}

	log.Printf("response of api call for %s: %v", sku, resp.GetQuantity())
}

func callUpdateQuantity(c api.InventoryClient) {
	log.Printf("calling service update")
	resp, err := c.UpdateQuantity(context.Background(), &api.UpdateQuantityRequest{
		Sku:      "Tablet",
		Quantity: 12,
	})

	if err != nil {
		log.Fatalf("Error when calc update: %v", err)
	}

	log.Printf("response of api call for Tablet update -2: %v:%v", resp.GetSku(), resp.GetQuantity())
}
