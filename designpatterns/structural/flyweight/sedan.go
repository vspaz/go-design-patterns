package flyweight

import "fmt"

type Sedan struct {
	carType string
}

func NewSedan(carType string) IVehicle {
	return &Sedan{
		carType: carType,
	}
}

func (t *Sedan) GetInfo(color string) string {
	return fmt.Sprintf("'%s' of color '%s' is created", t.carType, color)
}
