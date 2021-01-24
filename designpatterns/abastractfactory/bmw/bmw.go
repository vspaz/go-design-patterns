package bmw

import "go-design-patterns/designpatterns/abastractfactory"

type BMW struct {
}

func (b *BMW) MakeSedan() abastractfactory.ISedan {
	return &bmwSedan{
		abastractfactory.Sedan{
			Make:     "720i",
			PriceTag: 90000,
		},
	}
}

func (b *BMW) MakeSUV() abastractfactory.ISUV {
	return &bmwSUV{
		abastractfactory.SUV{
			Make:     "x5",
			PriceTag: 120000,
		},
	}
}
