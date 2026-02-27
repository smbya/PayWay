-- name: CreatePayment :one
INSERT INTO payments (user_id, amount, currency, status, destination, description)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, user_id, amount, currency, status, destination, description, created_at;

-- name: GetPaymentByID :one
SELECT id, user_id, amount, currency, status, destination, description, created_at
FROM payments
WHERE id = $1;