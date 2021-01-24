package bmw

import "go-design-patterns/designpatterns/abastractfactory"

type BMW struct {
}

func (b *BMW) makeSedan() abastractfactory.ISedan {
	return &bmwSedan{
		abastractfactory.Sedan{
			Make:     "720i",
			PriceTag: 90000,
		},
	}
}

func (b *BMW) makeSUV() abastractfactory.ISedan {
	return &bmwSUV{
		abastractfactory.SUV{
			Make:     "x5",
			PriceTag: 120000,
		},
	}
}
