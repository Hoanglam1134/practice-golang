-- name: CreateProduct :one
INSERT INTO inventory (
    sku,
    quantity
) VALUES (
    $1, $2
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM inventory
WHERE sku = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM inventory
ORDER BY sku;

-- name: UpdateProduct :one
UPDATE inventory
set quantity = $2
WHERE sku = $1
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM inventory
WHERE sku = $1;