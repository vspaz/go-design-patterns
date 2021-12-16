package template

type AbstractDrink interface {
}

type HotDrink struct {
	Drink AbstractDrink
}

func (h *HotDrink) BoilWater() string {
	return "water boiled"
}

func (h *HotDrink) PourInCup() string {
	return "poured in cup"
}
