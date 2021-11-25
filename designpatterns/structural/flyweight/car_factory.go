package flyweight

import (
	"log"
)

type CarFactory struct {
	nameToVehicle map[string]IVehicle
}

func (c *CarFactory) GetCar(carType string) IVehicle {
	car, _ := c.nameToVehicle[carType]
	if car != nil {
		return car
	}

	switch carType {
	case "sedan":
		car = NewSedan(carType)
	case "suv":
		car = NewSuv(carType)
	case "truck":
		car = NewTruck(carType)
	default:
		log.Fatalf("unknown type '%s' \n", carType)
	}
	return car
}
