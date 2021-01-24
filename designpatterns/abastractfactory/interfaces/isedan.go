package interfaces

type ISedan interface {
	SetMake(make string)
	GetMake() string

	SetPriceTag(price int)
	GetPriceTag() int
}

type Sedan struct {
	Make     string
	PriceTag int
}

func (s *Sedan) SetMake(make string) {
	s.Make = make
}

func (s *Sedan) GetMake() string {
	return s.Make
}

func (s *Sedan) SetPriceTag(price int) {
	s.PriceTag = price
}

func (s *Sedan) GetPriceTag() int {
	return s.PriceTag
}
