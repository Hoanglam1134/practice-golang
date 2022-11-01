package inventory

import (
	"context"
	"database/sql"
	"fmt"
)

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

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx Err: %v, rbErr: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

func (store *Store) CheckOrder(ctx context.Context, sku string, quantity int) (Inventory, error) {
	var product Inventory
	var errDB error
	err := store.execTx(ctx, func(q *Queries) error {
		product, errDB = q.GetProduct(ctx, sku)
		if errDB != nil {
			return errDB
		}

		if product.Quantity < int32(quantity) {
			return fmt.Errorf("Not enough quantity %d", product.Quantity)
		} else {
			arg := UpdateProductParams{
				Sku:      sku,
				Quantity: product.Quantity - int32(quantity),
			}

			_, errQueries := q.UpdateProduct(ctx, arg)
			if errQueries != nil {
				return errQueries
			}

		}
		return nil
	})
	return product, err
}
