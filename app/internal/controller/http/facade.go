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
	repo app.PaymentService
}

func NewPaymentFacade(repo app.PaymentService) PaymentFacade {
	return &facade{repo: repo}
}

func (f *facade) CreatePayment(c HttpContext) (string, int, error) {
	return f.repo.Create(context.TODO(), 2234)
}

func (f *facade) GetPaymentStatus(c HttpContext) (string, int, error) {
	return f.repo.GetStatus(context.TODO(), c.UrlParams["id"])
}
