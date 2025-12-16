package service

import "github.com/satyamdash/LLD-Golang/DesignPatterns/BookMyshow/model"

func IsSeatAvailable(show *model.Show, seatID string) bool {
	return !show.BookedSeatID[seatID]
}

func LockSeats(show *model.Show, seatIDs []string) bool {
	for _, seatID := range seatIDs {
		if show.BookedSeatID[seatID] {
			return false
		}
	}
	for _, seatID := range seatIDs {
		show.BookedSeatID[seatID] = true
	}
	return true
}
