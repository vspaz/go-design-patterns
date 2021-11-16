package adapter

import "fmt"

type Legacy interface {
	QuerySoap() string
}

type Soap struct{}

func (s *Soap) QuerySoap() string {
	return fmt.Sprintf("received response")
}
