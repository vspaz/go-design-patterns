package builder

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuilderOk(t *testing.T) {
	carBuilder := NewCar()
	car := carBuilder.
		WithChasis("aluminium").
		WithGasEngine(true).
		WithMake("bmw").
		Build()
	assert.False(t, car.Specs.ElectricalMotor)
	assert.True(t, car.Specs.GasEngine)
	assert.Equal(t, car.Make, "bmw")
}
