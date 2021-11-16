package adapter

import "fmt"

type Legacy interface {
	Decode(body string) string
}

type Soap struct {}

func (s *Soap) Decode(body string) string {
	return fmt.Sprintf("body decoded: %v", body)
}