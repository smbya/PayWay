package app

import "context"

type PaymentService interface {
	Create(ctx context.Context, amount float64) (string, int, error)
	GetStatus(ctx context.Context, id string) (string, int, error)
}

type paymentService struct {
	//repository
}

func NewPaymentService() PaymentService {
	return &paymentService{}
}

func (s *paymentService) Create(ctx context.Context, amount float64) (string, int, error) {
	return "create payment id 123", 200, nil
}

func (s *paymentService) GetStatus(ctx context.Context, id string) (string, int, error) {
	return "nu tipa status v rabote...", 200, nil
}
