package repository

import (
	"context"
	"log"
	"payway/internal/repository/posgresql/db/query"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) CreatePayment(ctx context.Context, user int, amount string, currency string, status string, destination string) string {
	q := query.New(r.pool)

	var am pgtype.Numeric

	am.Scan(amount)

	payment, err := q.CreatePayment(ctx, query.CreatePaymentParams{
		UserID:      int32(user),
		Amount:      am,
		Currency:    currency,
		Status:      status,
		Destination: destination,
	})

	if err != nil {
		log.Print("create payment error", err)
	}

	return payment.ID.String()

}
