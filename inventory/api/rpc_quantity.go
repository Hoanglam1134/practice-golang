package api

import (
	"context"
	"database/sql"
	"log"

	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

func (server *Server) GetQuantity(ctx context.Context, req *GetQuantityRequest) (*GetQuantityResponse, error) {
	sku := req.GetSku()
	product, err := server.store.GetProduct(ctx, sku)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("No rows in DB")
			return nil, status.Errorf(codes.Internal, "method GetQuantity not implemented")

		}
		log.Fatalf("error when get product %v", err)
	}
	resp := &GetQuantityResponse{
		Quantity: product.Quantity,
	}

	return resp, nil
}

func (server *Server) UpdateQuantity(ctx context.Context, req *UpdateQuantityRequest) (*UpdateQuantityResponse, error) {
	sku := req.GetSku()
	quantity := req.GetQuantity()

	product, err := server.store.CheckOrder(ctx, sku, int(quantity))

	if err != nil {
		log.Fatal("Error when queries")
	}

	resp := &UpdateQuantityResponse{
		Sku:      product.Sku,
		Quantity: product.Quantity,
	}
	return resp, nil
}
