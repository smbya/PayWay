package posgresql

import (
	"context"
	"fmt"
	"payway/internal/repository"
	"payway/internal/repository/posgresql/db/query"
	"payway/pkg/domain/payment"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
)

type paymentRepo struct {
	pool    *pgxpool.Pool
	queries *query.Queries
}

func NewPaymentRepo(pool *pgxpool.Pool) repository.PaymentRepository {
	return &paymentRepo{
		pool:    pool,
		queries: query.New(pool),
	}
}

func (r *paymentRepo) CreatePayment(ctx context.Context, params repository.CreatePaymentParams) (*payment.Payment, error) {
	var amount pgtype.Numeric
	if err := amount.Scan(params.Amount); err != nil {
		return nil, err
	}

	var description pgtype.Text
	if err := description.Scan(params.Description); err != nil {
		return nil, err
	}

	dbPayment, err := r.queries.CreatePayment(ctx, query.CreatePaymentParams{
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

func (r *paymentRepo) GetPaymentByID(ctx context.Context, id string) (*payment.Payment, error) {
	var uuid pgtype.UUID
	if err := uuid.Scan(id); err != nil {
		return nil, err
	}

	dbPayment, err := r.queries.GetPaymentByID(ctx, uuid)
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
