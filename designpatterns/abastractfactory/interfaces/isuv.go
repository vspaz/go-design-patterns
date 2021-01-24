package interfaces

type ISUV interface {
	setMake(make string)
	getMake() string

	setPriceTag(price int)
	getPriceTag() int
}

type SUV struct {
	Make     string
	PriceTag int
}

func (s *SUV) setMake(make string) {
	s.Make = make
}

func (s *SUV) getMake() string {
	return s.Make
}

func (s *SUV) setPriceTag(price int) {
	s.PriceTag = price
}

func (s *SUV) getPriceTag() int {
	return s.PriceTag
}
