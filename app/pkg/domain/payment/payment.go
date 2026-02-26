package payment

import "time"

type Payment struct {
	ID          int64
	Amount      string
	Currency    string
	Status      Status
	Description string
	CreatedAt   time.Time
}

func CreatePayment(
	id int64,
	amount string,
	currency string,
	status Status,
	description string,
	createdAt time.Time,
) *Payment {
	return &Payment{
		ID:          id,
		Amount:      amount,
		Currency:    currency,
		Status:      status,
		Description: description,
		CreatedAt:   createdAt,
	}
}
