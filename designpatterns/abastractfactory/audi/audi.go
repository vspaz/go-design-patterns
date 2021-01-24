package audi

import "go-design-patterns/designpatterns/abastractfactory"

type AUDI struct {
}

func (a *AUDI) MakeSedan() abastractfactory.ISedan {
	return &audiSedan{
		abastractfactory.Sedan{
			Make:     "A8",
			PriceTag: 70000,
		},
	}
}

func (a *AUDI) MakeSUV() abastractfactory.ISUV {
	return &AudiSUV{
		abastractfactory.SUV{
			Make:     "Q7",
			PriceTag: 900000,
		},
	}
}
