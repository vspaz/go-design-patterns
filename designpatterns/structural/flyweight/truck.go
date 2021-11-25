package flyweight

import "fmt"

type Truck struct {
	carType string
}

func NewTruck(carType string) IVehicle {
	return &Truck{
		carType: carType,
	}
}

func (t *Truck) GetInfo(color string) string {
	return fmt.Sprintf("'%s' of color '%s' is created", t.carType, color)
}
