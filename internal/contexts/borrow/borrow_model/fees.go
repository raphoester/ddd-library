package borrow_model

type Fee struct {
	days  int
	price Price
}

type Price struct {
	rounds int
	cents  int
}

func PriceFromFloat(price float64) Price {
	return Price{
		rounds: int(price),
		cents:  int(price*100) % 100,
	}
}

func NewPrice(rounds int, cents int) Price {
	return Price{
		rounds: 0,
		cents:  0,
	}
}
