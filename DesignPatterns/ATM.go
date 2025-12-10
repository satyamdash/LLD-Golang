package designpatterns

import "fmt"

// Handler interface
type DispenseHandler interface {
	SetNext(handler DispenseHandler)
	Dispense(amount int)
}

// Concrete Handler
type NoteDispenser struct {
	denomination int
	next         DispenseHandler
}

func NewNoteDispenser(denomination int) *NoteDispenser {
	return &NoteDispenser{denomination: denomination}
}

func (d *NoteDispenser) SetNext(handler DispenseHandler) {
	d.next = handler
}

func (d *NoteDispenser) Dispense(amount int) {
	if amount <= 0 {
		return
	}

	if amount >= d.denomination {
		numNotes := amount / d.denomination
		remaining := amount % d.denomination
		fmt.Printf("Dispensing %d notes of %d\n", numNotes, d.denomination)

		if remaining != 0 && d.next != nil {
			d.next.Dispense(remaining)
		} else if remaining != 0 {
			fmt.Printf("Cannot dispense remaining amount: %d\n", remaining)
		}
	} else {
		// Pass request to the next handler
		if d.next != nil {
			d.next.Dispense(amount)
		}
	}
}

// ATM builds the chain
type ATM struct {
	chain DispenseHandler
}

func NewATM() *ATM {
	// Create handlers
	h2000 := NewNoteDispenser(2000)
	h500 := NewNoteDispenser(500)
	h200 := NewNoteDispenser(200)
	h100 := NewNoteDispenser(100)

	// Construct chain: 2000 -> 500 -> 200 -> 100
	h2000.SetNext(h500)
	h500.SetNext(h200)
	h200.SetNext(h100)

	return &ATM{
		chain: h2000,
	}
}

func (atm *ATM) Withdraw(amount int) {
	fmt.Printf("\nRequest to withdraw %d\n", amount)
	atm.chain.Dispense(amount)
}

// MAIN
func CashATM() {
	atm := NewATM()

	atm.Withdraw(3700)
	atm.Withdraw(900)
	atm.Withdraw(1300)
	atm.Withdraw(125) // will show remaining un-dispensable amount
}
