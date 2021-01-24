package abastractfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCarAbstractFactoryOk(t *testing.T) {
	audiFactory, _ := getCarAbstractFactory("audi")
	audiSedan := audiFactory.MakeSedan()
	assert.Equal(t, "A8", audiSedan.GetMake())
	assert.Equal(t, 70000, audiSedan.GetPriceTag())
}

func TestGetCarAbstractFactoryNotImplementedType(t *testing.T) {
	audiFactory, err := getCarAbstractFactory("tesla")
	assert.Equal(t, nil, audiFactory)
	assert.NotEqual(t, nil, err)
}
