package bridgepattern

import "fmt"

type Window struct {
	printer Printer
}

func (w *Window) Print() {
	fmt.Println("Print request for mac")
	w.printer.PrintFile()
}

func (w *Window) setPrinter(printertype Printer) {
	w.printer = printertype
}
