package bridgepattern

import "fmt"

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Print request for mac")
	m.printer.PrintFile()
}

func (m *Mac) setPrinter(printertype Printer) {
	m.printer = printertype
}
