package abastractfactory

import (
	"fmt"
	"go-design-patterns/designpatterns/abastractfactory/audi"
	"go-design-patterns/designpatterns/abastractfactory/bmw"
	"go-design-patterns/designpatterns/abastractfactory/interfaces"
)

type ICarFactory interface {
	MakeSedan() interfaces.ISedan
	MakeSUV() interfaces.ISUV
}

func getCarAbstractFactory(manufacturer string) (ICarFactory, error) {
	switch manufacturer {
	case "audi":
		return &audi.AUDI{}, nil
	case "bmw":
		return &bmw.BMW{}, nil
	default:
		return nil, fmt.Errorf("invalid manufacturer type '%s'", manufacturer)
	}
}
