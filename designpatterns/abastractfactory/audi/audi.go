package audi

import "go-design-patterns/designpatterns/abastractfactory"

type AUDI struct {
}

func (a *AUDI) makeSedan() abastractfactory.ISedan {
	return AudiSedan{}
}

func (a *AUDI) makeSUV() abastractfactory.ISUV {
	return AudiSUV{}
}
