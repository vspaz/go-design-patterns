package abastractfactory

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	audiFactory, _ := getCarAbstractFactory("audi")
	audiSedan := audiFactory.MakeSedan()
	assert.Equal(t, "A8", audiSedan.GetMake())
	assert.Equal(t, 70000, audiSedan.GetPriceTag())
}
