package orders

import (
	"context"
	"database/sql"
)

type StoreMock interface {
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	GetOrder(ctx context.Context, id int32) (Order, error)
	UpdateOrder(ctx context.Context, arg UpdateOrderParams) (Order, error)
}

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)

// 	if err != nil {
// 		return err
// 	}

// 	q := New(tx)
// 	err = fn(q)
// 	if err != nil {
// 		if rbErr := tx.Rollback(); rbErr != nil {
// 			return fmt.Errorf("tx Err: %v, rbErr: %v", err, rbErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }

// func (store *Store) CreateOrderTx(ctx context.Context, sku string, quantity int) (Order, error) {
// 	var product Order
// 	var errDB error
// 	err := store.execTx(ctx, func(q *Queries) error {
// 		product, errDB = q.GetProduct(ctx, sku)
// 		if errDB != nil {
// 			return errDB
// 		}

// 		if product.Quantity < int32(quantity) {
// 			return fmt.Errorf("Not enough quantity %d", product.Quantity)
// 		} else {
// 			arg := UpdateProductParams{
// 				Sku:      sku,
// 				Quantity: product.Quantity - int32(quantity),
// 			}

// 			_, errQueries := q.UpdateProduct(ctx, arg)
// 			if errQueries != nil {
// 				return errQueries
// 			}

// 		}
// 		return nil
// 	})
// 	return product, err
// }
