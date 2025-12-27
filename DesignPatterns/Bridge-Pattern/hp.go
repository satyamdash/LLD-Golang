package bridgepattern

import "fmt"

type hp struct{}

func (h *hp) PrintFile() {
	fmt.Println("hp file printing")
}
