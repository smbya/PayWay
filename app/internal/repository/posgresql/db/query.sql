-- name: CreatePayment :one
INSERT INTO payments (user_id, amount, currency, status, destination)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;