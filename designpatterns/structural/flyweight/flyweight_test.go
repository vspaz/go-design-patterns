package flyweight

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFlyweightOk(t *testing.T) {
	carFactory := NewCarFactory()
	suv_1 := carFactory.GetCar("suv")
	suv_2 := carFactory.GetCar("suv")
	assert.Equal(t, suv_1, suv_2)
	assert.Equal(t, 1, carFactory.GetCarCount())

	truck_1 := carFactory.GetCar("truck")
	assert.NotEqual(t, suv_1, truck_1)
	assert.Equal(t, 2, carFactory.GetCarCount())
}
