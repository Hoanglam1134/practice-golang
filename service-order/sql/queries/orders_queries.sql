-- name: CreateOrder :one
INSERT INTO orders (
  id, sku, quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: ListOrder :many
SELECT * FROM orders
ORDER BY id;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE id = $1;

-- name: UpdateOrder :one
UPDATE orders
set sku = $2,
quantity = $3
WHERE id = $1
RETURNING *;