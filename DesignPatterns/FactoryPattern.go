package designpatterns

type EPizza interface {
	Cost() float64
}

type EMargherita struct{}

func (EMargherita) Cost() float64 {
	return 200
}

type EFarmhouse struct{}

func (EFarmhouse) Cost() float64 {
	return 350
}

type EVeggieDelight struct{}

func (EVeggieDelight) Cost() float64 {
	return 300
}

func NewPizza(name string) EPizza {
	switch name {
	case "Margherita":
		return EMargherita{}
	case "Farmhouse":
		return EFarmhouse{}
	case "VeggieDelight":
		return EVeggieDelight{}
	default:
		return nil

	}
}
