package payment

import "time"

type Payment struct {
	ID          string
	UserID      int32
	Amount      string
	Currency    string
	Status      Status
	Destination string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewPayment(
	id string,
	userID int32,
	amount string,
	currency string,
	status Status,
	destination string,
	description string,
	createdAt time.Time,
	updatedAt time.Time,
) *Payment {
	return &Payment{
		ID:          id,
		UserID:      userID,
		Amount:      amount,
		Currency:    currency,
		Status:      status,
		Destination: destination,
		Description: description,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}
}

func NewPaymentFromService(
	userID int32,
	amount string,
	currency string,
	status Status,
	destination string,
	description string,
) *Payment {
	return &Payment{
		UserID:      userID,
		Amount:      amount,
		Currency:    currency,
		Status:      status,
		Destination: destination,
		Description: description,
	}
}
