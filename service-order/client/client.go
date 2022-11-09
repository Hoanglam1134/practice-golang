package main

import (
	"context"
	"log"

	"practice-golang/service-order/api"

	"google.golang.org/grpc"
)

func main() {
	cc, err := grpc.Dial("localhost:10987", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	defer cc.Close()

	client := api.NewOrdersClient(cc)
	//callCreateOrder(client)
	callCompleteOrder(client)
	log.Printf("service client %v", client)
}

func callCreateOrder(c api.OrdersClient) {
	log.Printf("calling service get")
	resp, err := c.CreateOrder(context.Background(), &api.CreateOrderRequest{
		Id:       2,
		Sku:      "Laptop",
		Quantity: 3,
	})

	if err != nil {
		log.Fatalf("Error when call create order: %v", err)
	}

	log.Printf("response of api call for laptop: %v", resp.GetQuantity())
}

func callCompleteOrder(c api.OrdersClient) {
	log.Print("calling service update/complete order ...")
	resp, err := c.CompleteOrder(context.Background(), &api.CompleteOrderRequest{
		Id: 1,
	})

	if err != nil {
		log.Printf("Error when call complete order: %v", err)
	}

	log.Printf("response of api call for Tablet update -2: %v:%v", resp.GetSku(), resp.GetQuantity())
}
