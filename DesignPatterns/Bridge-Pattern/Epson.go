package bridgepattern

import "fmt"

type Epson struct{}

func (e *Epson) PrintFile() {
	fmt.Println("Epson file printing")
}
