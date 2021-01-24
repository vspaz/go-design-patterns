package audi

import (
	"go-design-patterns/designpatterns/abastractfactory/interfaces"
)

type AUDI struct {
}

func (a *AUDI) MakeSedan() interfaces.ISedan {
	return &audiSedan{
		interfaces.Sedan{
			Make:     "A8",
			PriceTag: 70000,
		},
	}
}

func (a *AUDI) MakeSUV() interfaces.ISUV {
	return &AudiSUV{
		interfaces.SUV{
			Make:     "Q7",
			PriceTag: 900000,
		},
	}
}
