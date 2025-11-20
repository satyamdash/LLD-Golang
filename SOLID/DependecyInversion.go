package solid

import "fmt"

//Bad EXAMPLE

type FileLogger struct{}

func (FileLogger) Log(msg string) {
	fmt.Println("Logging to file:", msg)
}

// type PaymentService struct {
// 	logger FileLogger
// }

func (p PaymentService) Process() {
	p.logger.Log("Processing payment")
}

//PaymentService is tightly coupled to FileLogger.

type GenericLogger interface {
	Log(data string)
}

type PaymentService struct {
	logger GenericLogger
}
