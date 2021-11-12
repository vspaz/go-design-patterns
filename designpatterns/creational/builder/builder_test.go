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
		Build()
	assert.False(t, car.Specs.ElectricalMotor)
	assert.True(t, car.Specs.GasEngine)
}
