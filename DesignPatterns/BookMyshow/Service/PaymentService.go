package service

import "github.com/satyamdash/LLD-Golang/DesignPatterns/BookMyshow/model"

type PaymentService struct{}

func (ps *PaymentService) MakePayment(booking *model.Booking) *model.Payment {
	return &model.Payment{
		ID:        "payment-1",
		BookingID: booking.ID,
		Status:    "SUCCESS",
	}
}
