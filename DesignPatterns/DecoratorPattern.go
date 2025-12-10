package designpatterns

type Pizza interface {
	Price() float64
}

type Margherita struct{}

func (Margherita) Price() float64 {
	return 200
}

type Cheese struct {
	pizza Pizza
}

func (c Cheese) Price() float64 {
	return c.pizza.Price() + 40
}

type Olives struct {
	pizza Pizza
}

func (o Olives) Price() float64 {
	return o.pizza.Price() + 30
}

type Jalapeno struct {
	pizza Pizza
}

func (j Jalapeno) Price() float64 {
	return j.pizza.Price() + 35
}
