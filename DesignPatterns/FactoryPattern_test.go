package designpatterns

import (
	"fmt"
	"testing"
)

func TestFactory(t *testing.T) {
	pizza := NewPizza("Margherita")
	fmt.Printf("Price of the Pizza %v\n", pizza.Cost())
}
