//go:generate mockery --name=ClientAdapter --case underscore
package adapter

import (
	"context"
	"log"
	"practice-golang/inventory/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewClientInventory() *client {
	cc, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("err while dial %v", err)
	}

	newClient := api.NewInventoryClient(cc)
	return &client{
		client: newClient,
	}
}

type ClientAdapter interface {
	GetQuantity(ctx context.Context, in *api.GetQuantityRequest, opts ...grpc.CallOption) (*api.GetQuantityResponse, error)
	UpdateQuantity(ctx context.Context, in *api.UpdateQuantityRequest, opts ...grpc.CallOption) (*api.UpdateQuantityResponse, error)
}

type client struct {
	client api.InventoryClient
}

func (c *client) GetQuantity(ctx context.Context, in *api.GetQuantityRequest, opts ...grpc.CallOption) (*api.GetQuantityResponse, error) {
	return c.client.GetQuantity(ctx, in)
}

func (c *client) UpdateQuantity(ctx context.Context, in *api.UpdateQuantityRequest, opts ...grpc.CallOption) (*api.UpdateQuantityResponse, error) {
	return c.client.UpdateQuantity(ctx, in)
}
