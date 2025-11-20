package solid

import "fmt"

//Bad Example

type Worker interface {
	Work()
}

type Human struct{}

func (Human) Work() {
	fmt.Println("Human is working")
}

type Robot struct{}

func (Robot) Work() {
	fmt.Println("Robot is working")
}

type LazyPerson struct{}

func (LazyPerson) Work() {
	panic("Lazy person is not working!")
}

// LazyPerson breaks the expectation of Work() by panicking.

type NonWorker interface {
	Eat()
}

func (LazyPerson) Eat() {
	fmt.Println("lazy person eats")
}
