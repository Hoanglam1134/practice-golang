package orders

import (
	"context"
	"fmt"
	orders "practice-golang/service-order/sqlc"
	"practice-golang/service-order/sqlc/mocks"
	"testing"
)

func TestCreateOrder(t *testing.T) {
	type args struct {
		ctx context.Context
		req *orders.CreateOrderParams
	}

	testcases := []struct {
		name       string
		args       args
		want       *orders.Order
		wantErr    bool
		wantErrVal error
	}{
		{
			name: "TestCreateOrder succeed",
			args: args{
				ctx: context.Background(),
				req: &orders.CreateOrderParams{
					ID:       1,
					Sku:      "Laptop",
					Quantity: 10,
				},
			},
			want: &orders.Order{
				ID:       2,
				Sku:      "Laptop",
				Quantity: 10,
			},
			wantErr: false,
		},
		{
			name: "TestCreateOrder fail",
			args: args{
				ctx: context.Background(),
				req: &orders.CreateOrderParams{
					ID:  1,
					Sku: "Laptop",
				},
			},
			want:       nil,
			wantErr:    true,
			wantErrVal: fmt.Errorf("Bad request"),
		},
	}

	for _, tc := range testcases {
		mockStore := &mocks.StoreMock{}
		t.Run(tc.name, func(t *testing.T) {
			switch tc.name {
			case "TestCreateOrder succeed":
				mockStore.On("CreateOrder", tc.args.ctx).
					Return(orders.Order{
						ID:       2,
						Sku:      "pv",
						Quantity: 10,
					}, nil)
			case "TestCreateOrder fail":
				mockStore.On("CreateOrder", tc.args.ctx, tc.args.req).
					Return(nil, fmt.Errorf("Bad request"))
			default:
			}
		})
	}

}
