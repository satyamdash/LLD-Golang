package model

import "time"

type Movie struct {
	ID       string
	Name     string
	Duration time.Duration
}

type Theatre struct {
	ID      string
	Address string
	City    string
	Screens []*Screen
}

type Screen struct {
	ID    string
	Seats []*Seat
}

type Seat struct {
	ID       string
	Row      int
	Category string
}

type Show struct {
	ID           string
	Movie        *Movie
	Screen       *Screen
	StartTime    time.Time
	BookedSeatID map[string]bool
}

type Booking struct {
	ID     string
	Show   *Show
	Seats  []*Seat
	Amount float64
}

type Payment struct {
	ID        string
	BookingID string
	Status    string
}
