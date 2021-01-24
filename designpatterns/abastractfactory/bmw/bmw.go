package bmw

import (
	"go-design-patterns/designpatterns/abastractfactory/interfaces"
)

type BMW struct {
}

func (b *BMW) MakeSedan() interfaces.ISedan {
	return &bmwSedan{
		interfaces.Sedan{
			Make:     "720i",
			PriceTag: 90000,
		},
	}
}

func (b *BMW) MakeSUV() interfaces.ISUV {
	return &bmwSUV{
		interfaces.SUV{
			Make:     "x5",
			PriceTag: 120000,
		},
	}
}
