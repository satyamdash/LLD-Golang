package service

import (
	"errors"

	"github.com/satyamdash/LLD-Golang/DesignPatterns/BookMyshow/model"
)

type BookingService struct{}

func (bs *BookingService) CreateBooking(show *model.Show, seats []*model.Seat) (*model.Booking, error) {
	seatIDs := []string{}
	for _, seat := range seats {
		seatIDs = append(seatIDs, seat.ID)
	}

	if !LockSeats(show, seatIDs) {
		return nil, errors.New("seat already booked")
	}

	return &model.Booking{
		ID:    "booking-1",
		Show:  show,
		Seats: seats,
	}, nil
}
