package adapterpattern

import "fmt"

type WindowAdapter struct {
	window *Window
}

func (wa *WindowAdapter) InsertIntoLightningPort() {
	fmt.Println("Insert Window into ligehtning")
	wa.window.InsertIntoUSB()
}
