package solid

import "fmt"

// Bad Example
type Vehicle interface {
	Drive()
	Fly()
	Sail()
}

type Car struct{}

// func (Car) Drive() {
// 	fmt.Println("Car driving")
// }

func (Car) Fly() {
	panic("Car can't fly!")
}

func (Car) Sail() {
	panic("Car can't sail!")
}

// Car is forced to implement methods it can’t support.

type Drive interface {
	Drive()
}

type Fly interface {
	Fly()
}

type Sail interface {
	Sail()
}

func (Car) Drive() {
	fmt.Println("Car can be driven")
}
