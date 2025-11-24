package designpatterns

// Bad COde
// type Order struct{}

// func (o Order) ApplyDiscount(customerType string, amount float64) float64 {
// 	if customerType == "regular" {
// 		return amount * 0.95
// 	}
// 	if customerType == "vip" {
// 		return amount * 0.90
// 	}
// 	if customerType == "black" {
// 		return amount * 0.70
// 	}
// 	return amount
// }

// Good Code
type CustomerStrategy interface {
	Discount(amount float64) float64
}

type RegularCustomer struct{}

func (RegularCustomer) Discount(amount float64) float64 {
	return amount * 0.95
}

type VipCustomer struct{}

func (VipCustomer) Discount(amount float64) float64 {
	return amount * 0.90
}

type BlackCustomer struct{}

func (BlackCustomer) Discount(amount float64) float64 {
	return amount * 0.70
}

type Order struct {
	customertype CustomerStrategy
}

func (o Order) ApplyDiscount(amount float64) float64 {
	return o.customertype.Discount(amount)
}
