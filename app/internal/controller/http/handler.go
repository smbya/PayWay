package http

import (
	"context"
	"log/slog"
	"payway/internal/app"
	"payway/pkg/domain/payment"
)

type CreatePaymentRequest struct {
	UserID      int32  `json:"user_id"`
	Amount      string `json:"amount"`
	Currency    string `json:"currency"`
	Destination string `json:"destination"`
	Description string `json:"description,omitempty"`
}

type PaymentHandler struct {
	paymentService app.PaymentService
	logger         *slog.Logger
}

func NewPaymentHandler(paymentService app.PaymentService, logger *slog.Logger) *PaymentHandler {
	return &PaymentHandler{
		paymentService: paymentService,
		logger:         logger,
	}
}

func (h *PaymentHandler) CreatePayment(ctx context.Context, req *CreatePaymentRequest) (*payment.Payment, error) {
	h.logger.Info("Create payment endpoint called")

	result, err := h.paymentService.Create(ctx, &app.CreatePaymentRequest{
		UserID:      req.UserID,
		Amount:      req.Amount,
		Currency:    req.Currency,
		Destination: req.Destination,
		Description: req.Description,
	})
	if err != nil {
		h.logger.Error("Error creating payment", "error", err)
		return nil, err
	}

	h.logger.Info("Create payment endpoint finished", "id", result.ID)
	return result, nil
}

func (h *PaymentHandler) GetPaymentByID(ctx context.Context, id string) (*payment.Payment, error) {
	h.logger.Info("Get payment status endpoint called", "id", id)

	result, err := h.paymentService.GetByID(ctx, id)
	if err != nil {
		h.logger.Error("Error getting payment status", "id", id, "error", err)
		return nil, err
	}

	h.logger.Info("Get payment status endpoint finished", "id", id)
	return result, nil
}
