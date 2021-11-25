package flyweight

import (
	"log"
)

type CarFactory struct {
	typeToVehicle map[string]IVehicle
}

func NewCarFactory() *CarFactory {
	return &CarFactory{
		typeToVehicle: make(map[string]IVehicle),
	}
}

func (c *CarFactory) GetCar(carType string) IVehicle {
	car, _ := c.typeToVehicle[carType]
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
	c.typeToVehicle[carType] = car
	return car
}

func (c *CarFactory) GetCarCount() int {
	return len(c.typeToVehicle)
}
