package flyweight

import "fmt"

type Suv struct {
	carType string
}

func NewSuv(carType string) IVehicle {
	return &Suv{
		carType: carType,
	}
}

func (t *Suv) GetInfo(color string) string {
	return fmt.Sprintf("'%s' of color '%s' is created", t.carType, color)
}
