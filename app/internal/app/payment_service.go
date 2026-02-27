package app

import (
	"context"
	"log/slog"
	"payway/internal/repository"
)

type PaymentService interface {
	Create(ctx context.Context, amount float64) (string, int, error)
	GetStatus(ctx context.Context, id string) (string, int, error)
}

type paymentService struct {
	repository *repository.Repository
	logger     *slog.Logger
}

func NewPaymentService(repository *repository.Repository, logger *slog.Logger) PaymentService {
	return &paymentService{
		repository: repository,
		logger:     logger,
	}
}

func (s *paymentService) Create(ctx context.Context, amount float64) (string, int, error) {
	s.logger.Info("Creating payment", "amount", amount)

	result := s.repository.CreatePayment(
		ctx,
		654,
		"456.78",
		"RUB",
		"New",
		"wallet123",
	)

	s.logger.Info("Payment created", "id", result)
	return "create payment id: " + result, 200, nil
}

func (s *paymentService) GetStatus(ctx context.Context, id string) (string, int, error) {
	s.logger.Info("Getting payment status", "id", id)

	return "nu tipa status v rabote...", 200, nil
}
