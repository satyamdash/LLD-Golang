package designpatterns

import (
	"fmt"
	"testing"
)

func TestDecorator(t *testing.T) {
	var fresh Pizza = Margherita{}
	fresh = Cheese{pizza: fresh}
	fresh = Olives{pizza: fresh}

	fmt.Printf("Prize of the pizza %f\n", fresh.Price())
}
