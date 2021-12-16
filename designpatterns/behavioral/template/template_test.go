package template

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemplate(t *testing.T) {
	tea := &Tea{}
	hotTea := &HotDrink{
		Drink: tea,
	}
	assert.Equal(t, "water boiled", hotTea.BoilWater())
	assert.Equal(t, "tea bag steepedlemon added", hotTea.Drink.PrepareRecipe())
	assert.Equal(t, "poured in cup", hotTea.PourInCup())
}
