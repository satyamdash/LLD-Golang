package elevator

type Direction int

const (
	IDLE Direction = iota
	UP
	DOWN
)

type Elevator struct {
	ID           int
	CurrentFloor int
	Direction    Direction
	RequestQueue []int // floors to visit
}

func (e *Elevator) AddRequest(floor int) {
	e.RequestQueue = append(e.RequestQueue, floor)
}

func (e *Elevator) Move() {
	if len(e.RequestQueue) == 0 {
		e.Direction = IDLE
		return
	}

	next := e.RequestQueue[0]
	e.RequestQueue = e.RequestQueue[1:]

	if next > e.CurrentFloor {
		e.Direction = UP
	} else if next < e.CurrentFloor {
		e.Direction = DOWN
	} else {
		e.Direction = IDLE
	}

	e.CurrentFloor = next
}

// func main() {
// 	ec := elevator.ElevatorController{
// 		Elevators: []*elevator.Elevator{
// 			{ID: 1, CurrentFloor: 0},
// 			{ID: 2, CurrentFloor: 5},
// 		},
// 	}

// 	e := ec.AssignElevator(3)
// 	fmt.Println("Elevator Assigned:", e.ID)

// 	e.AddRequest(7)
// 	ec.StepAll()
// 	ec.StepAll()
// }
