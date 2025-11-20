package solid

// Bad EXample
func Discount(customerType string, amount float64) float64 {
	if customerType == "regular" {
		return amount * 0.9
	}
	if customerType == "vip" {
		return amount * 0.8
	}
	return amount
}

// Correction
type DiscountStrategy interface {
	Apply(amount float64) float64
}

type RegularDiscount struct{}

func (RegularDiscount) Apply(amount float64) float64 {
	return amount * 0.9
}

type VipDiscount struct{}

func (VipDiscount) Apply(amount float64) float64 {
	return amount * 0.8
}

func Calculate(d DiscountStrategy, amount float64) float64 {
	return d.Apply(amount)
}
