package http

import (
	"context"
	"payway/internal/app"
)

type PaymentFacade interface {
	CreatePayment(c HttpContext) (string, int, error)
	GetPaymentStatus(c HttpContext) (string, int, error)
}

type facade struct {
	PaymentServicie app.PaymentService
}

func NewPaymentFacade(PaymentServicie app.PaymentService) PaymentFacade {
	return &facade{PaymentServicie: PaymentServicie}
}

func (f *facade) CreatePayment(c HttpContext) (string, int, error) {
	return f.PaymentServicie.Create(context.TODO(), 2234)
}

func (f *facade) GetPaymentStatus(c HttpContext) (string, int, error) {
	return f.PaymentServicie.GetStatus(context.TODO(), c.UrlParams["id"])
}
