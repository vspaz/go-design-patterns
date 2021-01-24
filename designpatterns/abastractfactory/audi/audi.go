package audi

import "go-design-patterns/designpatterns/abastractfactory"

type AUDI struct {
}

func (a *AUDI) makeSedan() abastractfactory.ISedan {
	return &audiSedan{
		abastractfactory.Sedan{
			Make:     "A8",
			PriceTag: 100500,
		},
	}
}

func (a *AUDI) makeSUV() abastractfactory.ISUV {
	return &AudiSUV{
		abastractfactory.SUV{
			Make:     "Q7",
			PriceTag: 300500,
		},
	}
}
