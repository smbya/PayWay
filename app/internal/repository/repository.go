package repository

import (
	"context"
	"fmt"
	"payway/internal/repository/posgresql/db/query"
	"payway/pkg/domain/payment"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	pool *pgxpool.Pool
}

type CreatePaymentParams struct {
	UserID      int32
	Amount      string
	Currency    string
	Status      string
	Destination string
	Description string
}

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}

func (r *Repository) CreatePayment(ctx context.Context, params CreatePaymentParams) (*payment.Payment, error) {
	q := query.New(r.pool)

	var amount pgtype.Numeric
	if err := amount.Scan(params.Amount); err != nil {
		return nil, err
	}

	var description pgtype.Text
	if err := description.Scan(params.Description); err != nil {
		return nil, err
	}

	dbPayment, err := q.CreatePayment(ctx, query.CreatePaymentParams{
		UserID:      params.UserID,
		Amount:      amount,
		Currency:    params.Currency,
		Status:      params.Status,
		Destination: params.Destination,
		Description: description,
	})
	if err != nil {
		return nil, err
	}

	return toDomainPayment(dbPayment), nil
}

func (r *Repository) GetPaymentByID(ctx context.Context, id string) (*payment.Payment, error) {
	q := query.New(r.pool)

	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return nil, err
	}

	dbPayment, err := q.GetPaymentByID(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return toDomainPaymentFromGet(dbPayment), nil
}

func toDomainPayment(p query.CreatePaymentRow) *payment.Payment {
	amount, _ := p.Amount.Float64Value()
	return payment.NewPayment(
		p.ID.String(),
		p.UserID,
		formatAmount(amount.Float64),
		p.Currency,
		payment.Status(p.Status),
		p.Destination,
		p.Description.String,
		p.CreatedAt.Time,
		p.CreatedAt.Time,
	)
}

func toDomainPaymentFromGet(p query.GetPaymentByIDRow) *payment.Payment {
	amount, _ := p.Amount.Float64Value()
	return payment.NewPayment(
		p.ID.String(),
		p.UserID,
		formatAmount(amount.Float64),
		p.Currency,
		payment.Status(p.Status),
		p.Destination,
		p.Description.String,
		p.CreatedAt.Time,
		p.CreatedAt.Time,
	)
}

func formatAmount(amount float64) string {
	return fmt.Sprintf("%.2f", amount)
}
