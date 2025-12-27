package designpatterns

import (
	"fmt"
)

// Bad Code
type Order2 struct {
	status string
}

func (o *Order2) UpdateStatus(newStatus string) {
	o.status = newStatus

	fmt.Println("Sending SMS:", newStatus)
	fmt.Println("Sending Email:", newStatus)
	fmt.Println("Sending Push Notification:", newStatus)
}

// GGOOD CODE

type Observer interface {
	Status(status string)
}

type Notification struct {
	observers []Observer
	status    string
}

func (n *Notification) AddObserver(obs Observer) {
	n.observers = append(n.observers, obs)
}

func (n *Notification) RemoveObserver(obs Observer) {
	for i, o := range n.observers {
		if o == obs {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *Notification) SetStatus(status string) {
	n.status = status
	for _, o := range n.observers {
		o.Status(status)
	}
}

type SMS struct{}

func (SMS) Status(status string) {
	fmt.Printf("SMS Notification %v", status)
	fmt.Println()
}

type Email struct{}

func (Email) Status(status string) {
	fmt.Printf("EmailMS Notification %v", status)
	fmt.Println()
}

type PushNotification struct{}

func (PushNotification) Status(status string) {
	fmt.Printf("PushNotification Notification %v", status)
	fmt.Println()
}
