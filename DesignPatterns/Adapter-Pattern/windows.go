package adapterpattern

import "fmt"

type Window struct{}

func (w *Window) InsertIntoUSB() {
	fmt.Println("Insert Into USB")
}
