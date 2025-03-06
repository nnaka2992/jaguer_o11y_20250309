-- name: GetUserByID :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (
  id, name, email, password
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateAuthor :exec
UPDATE users
SET name = $2,
    email = $3,
    password = $4,
    updated_at = now()
WHERE id = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: GetItems :many
SELECT * FROM items LIMIT $1 OFFSET $2;

-- name: GetItemByID :one
SELECT * FROM items
WHERE id = $1 LIMIT 1;

-- name: CreateItem :one
INSERT INTO items (
  name, description, price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateItem :exec
UPDATE items
SET name = $2,
    description = $3,
    price = $4,
    updated_at = now()
WHERE id = $1;

-- name: DeleteItem :exec
DELETE FROM items
WHERE id = $1;

-- name: GetOrders :many
SELECT * FROM orders LIMIT $1 OFFSET $2;

-- name: GetOrderByID :one
SELECT * FROM orders
WHERE id = $1 LIMIT 1;

-- name: GetOrdersByUserID :many
SELECT * FROM orders
WHERE user_id = $1 LIMIT $2 OFFSET $3;

-- name: CreateOrder :one
INSERT INTO orders (
  user_id, item_id, quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteOrder :one
INSERT INTO orders (
  user_id, item_id, quantity
) VALUES (
  $1, $2, $3
)
RETURNING *;
