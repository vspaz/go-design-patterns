package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTeaOk(t *testing.T) {
	tea := &Tea{}
	hotTea := &HotDrink{
		Drink: tea,
	}
	assert.Equal(t, "water boiled", hotTea.BoilWater())
	assert.Equal(t, "tea bag steeped; lemon added", hotTea.PrepareRecipe())
	assert.Equal(t, "poured in cup", hotTea.PourInCup())
}

func TestCoffeeOk(t *testing.T) {
	hotCoffee := &Coffee{}
	hotDrink := &HotDrink{
		Drink: hotCoffee,
	}
	assert.Equal(t, "water boiled", hotDrink.BoilWater())
	assert.Equal(t, "coffee brewed; sugar and milk added to coffee", hotDrink.PrepareRecipe())
	assert.Equal(t, "poured in cup", hotDrink.PourInCup())
}
