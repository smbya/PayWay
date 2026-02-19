package app

import (
	"context"
	"payway/internal/repository"
)

type PaymentService interface {
	Create(ctx context.Context, amount float64) (string, int, error)
	GetStatus(ctx context.Context, id string) (string, int, error)
}

type paymentService struct {
	repository *repository.Repository
}

func NewPaymentService(repository *repository.Repository) PaymentService {
	return &paymentService{
		repository: repository,
	}
}

func (s *paymentService) Create(ctx context.Context, amount float64) (string, int, error) {

	result := s.repository.CreatePayment(
		ctx,
		654,
		"456.78",
		"RUB",
		"New",
		"wallet123",
	)

	return "create payment id: " + result, 200, nil
}

func (s *paymentService) GetStatus(ctx context.Context, id string) (string, int, error) {
	return "nu tipa status v rabote...", 200, nil
}
