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
		WithSeats(6).
		Build()
	assert.False(t, car.Specs.ElectricalMotor)
	assert.True(t, car.Specs.GasEngine)
	assert.Equal(t, car.Specs.Seats, 6)
	assert.Equal(t, car.Make, "bmw")
}
