package designpatterns

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// --- Types & Enums ---

type VehicleType int

const (
	Motorcycle VehicleType = iota
	Car
	Bus
)

type SpotSize int

const (
	Small SpotSize = iota
	Medium
	Large
)

// --- Domain Models ---

type Vehicle struct {
	Plate string
	Type  VehicleType
}

type ParkingSpot struct {
	ID       string
	Size     SpotSize
	LevelID  int
	Occupied bool
	// occupiedBy TicketID or vehicle plate could be stored
}

type Ticket struct {
	ID        string
	Vehicle   Vehicle
	SpotID    string
	LevelID   int
	EntryTime time.Time
	ExitTime  *time.Time
	Paid      bool
	Fee       float64
}

// --- Interfaces ---

// Allocator chooses a free spot for a vehicle
type Allocator interface {
	FindSpot(l *ParkingLot, v Vehicle) (*ParkingSpot, error)
}

// PaymentCalculator calculates fee given entry/exit times and vehicle type
type PaymentCalculator interface {
	Calculate(v Vehicle, entry, exit time.Time) float64
}

// --- Implementations ---

// SimpleFirstFitAllocator: scans levels and spots, returns first fit.
type SimpleFirstFitAllocator struct{}

func (a *SimpleFirstFitAllocator) FindSpot(l *ParkingLot, v Vehicle) (*ParkingSpot, error) {
	// Lock provided externally by ParkingLot methods
	for _, level := range l.Levels {
		for i := range level.Spots {
			spot := &level.Spots[i]
			if !spot.Occupied && canFit(spot.Size, v.Type) {
				return spot, nil
			}
		}
	}
	return nil, errors.New("no spot available")
}

func canFit(size SpotSize, vt VehicleType) bool {
	switch vt {
	case Motorcycle:
		return true // fits anywhere
	case Car:
		return size == Medium || size == Large
	case Bus:
		return size == Large
	default:
		return false
	}
}

// HourlyPaymentCalculator: simple per-hour rates
type HourlyPaymentCalculator struct {
	Rates map[VehicleType]float64 // per hour
}

func (pc *HourlyPaymentCalculator) Calculate(v Vehicle, entry, exit time.Time) float64 {
	duration := exit.Sub(entry)
	hours := duration.Hours()
	if hours < 1 {
		hours = 1
	}
	rate, ok := pc.Rates[v.Type]
	if !ok {
		rate = 10 // fallback
	}
	return rate * hours
}

// --- Parking Level & Lot ---

type Level struct {
	ID    int
	Spots []ParkingSpot
}

type ParkingLot struct {
	sync.Mutex
	ID                 string
	Levels             []*Level
	tickets            map[string]*Ticket
	spotIndex          map[string]*ParkingSpot // quick lookup by spot id
	allocator          Allocator
	paymentCalculator  PaymentCalculator
	nextTicketSequence int
}

// NewParkingLot creates a ParkingLot with given levels
func NewParkingLot(id string, levels []*Level, alloc Allocator, pc PaymentCalculator) *ParkingLot {
	spotIndex := make(map[string]*ParkingSpot)
	for _, lvl := range levels {
		for i := range lvl.Spots {
			s := &lvl.Spots[i]
			spotIndex[s.ID] = s
		}
	}
	return &ParkingLot{
		ID:                id,
		Levels:            levels,
		tickets:           make(map[string]*Ticket),
		spotIndex:         spotIndex,
		allocator:         alloc,
		paymentCalculator: pc,
	}
}

// Entry: park the vehicle
func (pl *ParkingLot) Enter(v Vehicle) (*Ticket, error) {
	pl.Lock()
	defer pl.Unlock()

	spot, err := pl.allocator.FindSpot(pl, v)
	if err != nil {
		return nil, err
	}
	// mark occupied
	spot.Occupied = true
	pl.nextTicketSequence++
	tid := fmt.Sprintf("%s-T%06d", pl.ID, pl.nextTicketSequence)
	ticket := &Ticket{
		ID:        tid,
		Vehicle:   v,
		SpotID:    spot.ID,
		LevelID:   spot.LevelID,
		EntryTime: time.Now(),
	}
	pl.tickets[tid] = ticket
	return ticket, nil
}

// Exit: unpark the vehicle and compute fee
func (pl *ParkingLot) Exit(ticketID string) (*Ticket, error) {
	pl.Lock()
	defer pl.Unlock()

	t, ok := pl.tickets[ticketID]
	if !ok {
		return nil, errors.New("invalid ticket")
	}
	if t.ExitTime != nil {
		return nil, errors.New("ticket already used to exit")
	}
	now := time.Now()
	t.ExitTime = &now
	fee := pl.paymentCalculator.Calculate(t.Vehicle, t.EntryTime, now)
	t.Fee = fee
	t.Paid = true // assume paid for simplicity
	// free spot
	spot, ok := pl.spotIndex[t.SpotID]
	if ok {
		spot.Occupied = false
	}
	return t, nil
}

// Query available spots per level / overall
func (pl *ParkingLot) AvailableSpots() map[int]int {
	pl.Lock()
	defer pl.Unlock()
	res := map[int]int{}
	for _, lvl := range pl.Levels {
		cnt := 0
		for i := range lvl.Spots {
			if !lvl.Spots[i].Occupied {
				cnt++
			}
		}
		res[lvl.ID] = cnt
	}
	return res
}
