package server

import (
	"context"
	"errors"
	"log"
	inventoryApi "practice-golang/inventory/api"
	"practice-golang/service-order/api"
	orders "practice-golang/service-order/sqlc"
)

func (server *Server) CreateOrder(ctx context.Context, req *api.CreateOrderRequest) (*api.CreateOrderResponse, error) {
	product, err := server.store.CreateOrder(ctx, orders.CreateOrderParams{
		ID:       req.GetId(),
		Sku:      req.GetSku(),
		Quantity: req.GetQuantity(),
	})
	if err != nil {
		log.Printf("error when create product %v", err)
		return nil, err
	}
	resp := &api.CreateOrderResponse{
		Id:       product.ID,
		Sku:      product.Sku,
		Quantity: product.Quantity,
	}

	return resp, nil
}

func (server *Server) CompleteOrder(ctx context.Context, req *api.CompleteOrderRequest) (*api.CompleteOrderResponse, error) {
	id := req.GetId()
	order, err := server.store.GetOrder(ctx, id)
	if err != nil {
		log.Print("error when query get order")
		return nil, err
	}
	log.Printf("order after query: %s : %d", order.Sku, order.Quantity)

	// check ton kho ben Inventory Service
	product, errClient := server.inventoryClient.GetQuantity(ctx, &inventoryApi.GetQuantityRequest{
		Sku: order.Sku,
	})
	if errClient != nil {
		log.Fatalf("Error when call api from inventory service: %v", errClient)
	}

	// Kiem tra du ton, neu khong du thi return nil
	if product.Quantity < order.Quantity {
		return nil, errors.New("Khong du ton kho!")
	}

	// if the quantity of product is still enough, decrease the quantity
	_, errClient = server.inventoryClient.UpdateQuantity(ctx, &inventoryApi.UpdateQuantityRequest{
		Sku:      order.Sku,
		Quantity: order.Quantity,
	})

	if errClient != nil {
		log.Print("Error when update quantity from inventory service")
		return nil, errClient
	}

	resp := &api.CompleteOrderResponse{
		Id:       order.ID,
		Sku:      order.Sku,
		Quantity: order.Quantity,
	}
	return resp, nil
}
