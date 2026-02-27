package repository

import (
	"context"
	"payway/pkg/domain/payment"
)

type CreatePaymentParams struct {
	UserID      int32
	Amount      string
	Currency    string
	Status      string
	Destination string
	Description string
}

type PaymentRepository interface {
	CreatePayment(ctx context.Context, params CreatePaymentParams) (*payment.Payment, error)
	GetPaymentByID(ctx context.Context, id string) (*payment.Payment, error)
}
