package flyweight

type Sedan struct {
	carType string
}

func NewSedan(carType string) *Sedan {
	return &Sedan{
		carType: carType,
	}
}
