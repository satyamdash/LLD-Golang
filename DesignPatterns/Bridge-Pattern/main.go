package bridgepattern

import "fmt"

func main() {

	hpPrinter := &hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.setPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.setPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &Window{}

	winComputer.setPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.setPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}
