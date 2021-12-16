package template

type Coffee struct{}

func (c *Coffee) PrepareRecipe() string {
	return c.BrewCoffee() + c.AddSugarAndMilk()
}

func (c *Coffee) BrewCoffee() string {
	return "coffee brewed; "
}

func (c *Coffee) AddSugarAndMilk() string {
	return "sugar and milk added to coffee"
}
