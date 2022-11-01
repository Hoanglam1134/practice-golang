package inventory

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMakeOrderTx(t *testing.T) {
	store := NewStore(testDB)
	product1 := RandomProduct(t)
	product2 := RandomProduct(t)

	errs := make(chan error)
	results := make(chan Inventory)

	go func() {
		result, err := store.CheckOrder(context.Background(), product1.Sku, int(product1.Quantity)-2)

		errs <- err
		results <- result
	}()

	go func() {
		result, err := store.CheckOrder(context.Background(), product2.Sku, int(product2.Quantity)+1)

		require.Error(t, err)
		results <- result

	}()

	err := <-errs
	require.NoError(t, err)

}
