package elevator

import (
	"math"
)

type ElevatorController struct {
	Elevators []*Elevator
}

func (c *ElevatorController) AssignElevator(requestFloor int) *Elevator {
	best := c.Elevators[0]
	minDist := math.MaxInt32

	for _, e := range c.Elevators {
		dist := int(math.Abs(float64(e.CurrentFloor - requestFloor)))
		if dist < minDist {
			minDist = dist
			best = e
		}
	}

	best.AddRequest(requestFloor)
	return best
}

func (c *ElevatorController) StepAll() {
	for _, e := range c.Elevators {
		e.Move()
	}
}
