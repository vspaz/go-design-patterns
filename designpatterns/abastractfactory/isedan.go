package abastractfactory

type ISedan interface {
	setMake(make string)
	getMake() string

	setPriceTag(price string)
	getPriceTag() int
}

type Sedan struct {
	Make     string
	PriceTag int
}

func (s *Sedan) setMake(make string) {
	s.Make = make
}

func (s *Sedan) getMake() string {
	return s.Make
}

func (s *Sedan) setPriceTag(price int) {
	s.PriceTag = price
}

func (s *Sedan) getPriceTag() int {
	return s.PriceTag
}
