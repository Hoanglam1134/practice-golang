-- name: CreateProduct :one
INSERT INTO inventory (
    sku,
    quantity
) VALUES (
    $1, $2
)
RETURNING *;