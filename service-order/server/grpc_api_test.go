package server

import (
	"context"
	"errors"
	inventoryApi "practice-golang/inventory/api"
	mockClient "practice-golang/service-order/adapter/mocks"
	"practice-golang/service-order/api"
	orders "practice-golang/service-order/sqlc"
	mockStore "practice-golang/service-order/sqlc/mocks"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestCreateOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.CreateOrderRequest
	}

	testcases := []struct {
		name       string
		args       args
		want       *api.CreateOrderResponse
		wantErr    bool
		wantErrVal error
	}{
		{
			name: "TestCreateOrder succeed",
			args: args{
				ctx: context.Background(),
				req: &api.CreateOrderRequest{
					Id:       1,
					Sku:      "Laptop",
					Quantity: 10,
				},
			},
			want: &api.CreateOrderResponse{
				Id:       1,
				Sku:      "Laptop",
				Quantity: 10,
			},
			wantErr: false,
		},
		{
			name: "TestCreateOrder fail",
			args: args{
				ctx: context.Background(),
				req: &api.CreateOrderRequest{
					Id:       1,
					Sku:      "Laptop",
					Quantity: 10,
				},
			},
			want:       nil,
			wantErr:    true,
			wantErrVal: errors.New("fail"),
		},
	}

	for _, tc := range testcases {
		mockStore := &mockStore.StoreQuerier{}
		server := Server{
			store: mockStore,
		}
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "TestCreateOrder succeed":
				mockStore.On("CreateOrder", tc.args.ctx, orders.CreateOrderParams{
					ID:       1,
					Sku:      "Laptop",
					Quantity: 10,
				}).
					Return(orders.Order{
						ID:       1,
						Sku:      "Laptop",
						Quantity: 10,
					}, nil)
			case "TestCreateOrder fail":
				mockStore.On("CreateOrder", tc.args.ctx, orders.CreateOrderParams{
					ID:       1,
					Sku:      "Laptop",
					Quantity: 10,
				}).
					Return(orders.Order{}, errors.New("fail"))
			default:
			}

			result, err := server.CreateOrder(tc.args.ctx, tc.args.req)
			require.Equal(t, tc.want, result)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}

}

func TestCompleteOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		req *api.CompleteOrderRequest
	}
	testcases := []struct {
		name       string
		args       args
		want       *api.CompleteOrderResponse
		wantErr    bool
		wantErrVal error
	}{
		{
			name: "TestCompleteOrder success, enough quantity",
			args: args{
				ctx: context.Background(),
				req: &api.CompleteOrderRequest{
					Id: 1,
				},
			},
			want: &api.CompleteOrderResponse{
				Id:       1,
				Sku:      "Laptop",
				Quantity: 3,
			},
			wantErr: false,
		},
		{
			name: "TestCompleteOrder fail, not enough quantity",
			args: args{
				ctx: context.Background(),
				req: &api.CompleteOrderRequest{
					Id: 1,
				},
			},
			want:       nil,
			wantErr:    true,
			wantErrVal: errors.New("not enough quantity!"),
		},
	}

	for _, tc := range testcases {
		mockStore := &mockStore.StoreQuerier{}
		mockClient := &mockClient.ClientAdapter{}
		server := Server{
			store:           mockStore,
			inventoryClient: mockClient,
		}
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "TestCompleteOrder success, enough quantity":
				mockStore.On("GetOrder", tc.args.ctx, mock.Anything).
					Return(orders.Order{
						ID:       1,
						Sku:      "Laptop",
						Quantity: 3,
					}, nil)
				mockClient.On("GetQuantity", tc.args.ctx, &inventoryApi.GetQuantityRequest{
					Sku: "Laptop",
				}).
					Return(&inventoryApi.GetQuantityResponse{
						Quantity: 10,
					}, nil)
				mockClient.On("UpdateQuantity", tc.args.ctx, mock.Anything).
					Return(&inventoryApi.UpdateQuantityResponse{
						Sku:      "Laptop",
						Quantity: 7,
					}, nil)
			case "TestCompleteOrder fail, not enough quantity":
				mockStore.On("GetOrder", tc.args.ctx, mock.Anything).
					Return(orders.Order{
						ID:       1,
						Sku:      "Laptop",
						Quantity: 5,
					}, nil)
				mockClient.On("GetQuantity", tc.args.ctx, mock.Anything).
					Return(&inventoryApi.GetQuantityResponse{
						Quantity: 3,
					}, nil)
			default:
			}

			result, err := server.CompleteOrder(tc.args.ctx, tc.args.req)
			require.Equal(t, tc.want, result)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
