package app

import (
	"context"
	"log/slog"
	"payway/internal/repository"
	"payway/pkg/domain/payment"
)

type PaymentService interface {
	Create(ctx context.Context, req *CreatePaymentRequest) (*payment.Payment, error)
	GetByID(ctx context.Context, id string) (*payment.Payment, error)
}

type CreatePaymentRequest struct {
	UserID      int32
	Amount      string
	Currency    string
	Destination string
	Description string
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

func (s *paymentService) Create(ctx context.Context, req *CreatePaymentRequest) (*payment.Payment, error) {
	s.logger.Info("Creating payment", "user_id", req.UserID, "amount", req.Amount)

	p, err := s.repository.CreatePayment(ctx, repository.CreatePaymentParams{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Currency:    req.Currency,
		Status:      string(payment.STATUS_NEW),
		Destination: req.Destination,
		Description: req.Description,
	})
	if err != nil {
		s.logger.Error("Failed to create payment", "error", err)
		return nil, err
	}

	s.logger.Info("Payment created", "id", p.ID)
	return p, nil
}

func (s *paymentService) GetByID(ctx context.Context, id string) (*payment.Payment, error) {
	s.logger.Info("Getting payment status", "id", id)

	p, err := s.repository.GetPaymentByID(ctx, id)
	if err != nil {
		s.logger.Error("Failed to get payment", "id", id, "error", err)
		return nil, err
	}

	return p, nil
}
