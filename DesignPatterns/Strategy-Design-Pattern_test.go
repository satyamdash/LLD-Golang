package designpatterns

import (
	"fmt"
	"testing"
)

func TestStrategyfunction(t *testing.T) {
	order := &Order{RegularCustomer{}}
	fmt.Printf("Regular Customer amount to pay after Discount %v \n", order.ApplyDiscount(100))

	order = &Order{BlackCustomer{}}
	fmt.Printf("Regular Customer amount to pay after Discount %v\n", order.ApplyDiscount(100))

	order = &Order{VipCustomer{}}
	fmt.Printf("VIP Customer amount to pay after Discount %v\n", order.ApplyDiscount(100))

}
