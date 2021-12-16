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
