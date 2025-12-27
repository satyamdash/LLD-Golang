package bridgepattern

type Computer interface {
	Print()
	setPrinter(Printer)
}
