package adapter

type Legacy interface {
	QuerySoap() string
}

type Soap struct{}

func (s *Soap) QuerySoap() string {
	return "making SOAP requests"
}
