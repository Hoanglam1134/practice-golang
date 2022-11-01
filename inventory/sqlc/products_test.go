package inventory

import (
	"context"
	"database/sql"
	"testing"
	"todolist-grpc/inventory/utils"

	"github.com/stretchr/testify/require"
)

func TestCreateHardProduct(t *testing.T) {
	arg := CreateProductParams{
		Sku:      "Watch",
		Quantity: 15,
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Sku, product.Sku)
	require.Equal(t, arg.Quantity, product.Quantity)
}
func RandomProduct(t *testing.T) Inventory {
	arg := CreateProductParams{
		Sku:      utils.RandomSku(),
		Quantity: int32(utils.RandomQuantity()),
	}

	product, err := testQueries.CreateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product)

	require.Equal(t, arg.Sku, product.Sku)
	require.Equal(t, arg.Quantity, product.Quantity)
	return product
}
func TestCreateProducts(t *testing.T) {
	RandomProduct(t)
}

func TestGetProduct(t *testing.T) {
	product1 := RandomProduct(t)
	product2, err := testQueries.GetProduct(context.Background(), product1.Sku)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Sku, product2.Sku)
	require.Equal(t, product1.Quantity, product2.Quantity)
}

func TestListProduct(t *testing.T) {
	for i := 0; i < 10; i++ {
		RandomProduct(t)
	}

	products, err := testQueries.ListProducts(context.Background())
	require.NoError(t, err)
	for _, product := range products {
		require.NotEmpty(t, product)
	}
}

// hoi thac mac cho nay can giai dap
func TestUpdateProduct(t *testing.T) {
	product1 := RandomProduct(t)
	arg := UpdateProductParams{
		Sku:      product1.Sku,
		Quantity: int32(utils.RandomQuantity()),
	}
	product2, err := testQueries.UpdateProduct(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, product2)

	require.Equal(t, product1.Sku, product2.Sku)
	require.Equal(t, arg.Quantity, product2.Quantity)
}

func TestDeleteProduct(t *testing.T) {
	product1 := RandomProduct(t)
	err := testQueries.DeleteProduct(context.Background(), product1.Sku)
	require.NoError(t, err)

	product2, err := testQueries.GetProduct(context.Background(), product1.Sku)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, product2)
}
