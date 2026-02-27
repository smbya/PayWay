package http

import (
	"context"
	"log/slog"
	"payway/internal/app"
)

type PaymentFacade interface {
	CreatePayment(c HttpContext) (string, int, error)
	GetPaymentStatus(c HttpContext) (string, int, error)
}

type facade struct {
	PaymentServicie app.PaymentService
	logger          *slog.Logger
}

func NewPaymentFacade(PaymentServicie app.PaymentService, logger *slog.Logger) PaymentFacade {
	return &facade{
		PaymentServicie: PaymentServicie,
		logger:          logger,
	}
}

func (f *facade) CreatePayment(c HttpContext) (string, int, error) {
	f.logger.Info("Create payment endpoint called")

	result, status, err := f.PaymentServicie.Create(context.TODO(), 2234)
	if err != nil {
		f.logger.Error("Error creating payment", "error", err)
	}

	f.logger.Info("Create payment endpoint finished", "status", status)

	return result, status, err
}

func (f *facade) GetPaymentStatus(c HttpContext) (string, int, error) {
	f.logger.Info("Get payment status endpoint called", "id", c.UrlParams["id"])

	result, status, err := f.PaymentServicie.GetStatus(context.TODO(), c.UrlParams["id"])
	if err != nil {
		f.logger.Error("Error getting payment status", "id", c.UrlParams["id"], "error", err)
	}

	f.logger.Info("Get payment status endpoint finished", "id", c.UrlParams["id"], "status", status)

	return result, status, err
}
