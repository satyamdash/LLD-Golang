package designpatterns

import "fmt"

// Log Levels
const (
	INFO = iota
	DEBUG
	ERROR
)

// Handler interface
type LogHandler interface {
	SetNext(next LogHandler)
	Handle(msg string, level int)
}

// Base handler with next reference
type BaseHandler struct {
	next LogHandler
}

func (b *BaseHandler) SetNext(next LogHandler) {
	b.next = next
}

func (b *BaseHandler) CallNext(msg string, level int) {
	if b.next != nil {
		b.next.Handle(msg, level)
	}
}

// ------------ Concrete Handlers ------------

// Info Logs
type InfoHandler struct {
	BaseHandler
}

func (i *InfoHandler) Handle(msg string, level int) {
	if level == INFO {
		fmt.Println("[INFO] :", msg)
	}
	i.CallNext(msg, level)
}

// Debug Logs
type DebugHandler struct {
	BaseHandler
}

func (d *DebugHandler) Handle(msg string, level int) {
	if level == DEBUG {
		fmt.Println("[DEBUG]:", msg)
	}
	d.CallNext(msg, level)
}

// Error Logs
type ErrorHandler struct {
	BaseHandler
}

func (e *ErrorHandler) Handle(msg string, level int) {
	if level == ERROR {
		fmt.Println("[ERROR]:", msg)
	}
	e.CallNext(msg, level)
}

// ------------ Client Logger ------------

type Logger struct {
	handler LogHandler
}

func NewLogger() *Logger {
	// Build chain: Info → Debug → Error
	info := &InfoHandler{}
	debug := &DebugHandler{}
	err := &ErrorHandler{}

	info.SetNext(debug)
	debug.SetNext(err)

	return &Logger{handler: info}
}

func (l *Logger) Log(msg string, level int) {
	l.handler.Handle(msg, level)
}

// ------------ MAIN ------------

func LoggerChain() {
	logger := NewLogger()

	logger.Log("App started", INFO)
	logger.Log("Loaded config file", DEBUG)
	logger.Log("Null pointer exception!", ERROR)
}
