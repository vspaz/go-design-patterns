package template

type Tea struct {
}

func (t *Tea) PrepareRecipe() string {
	return t.SteepTeaBag() + t.AddLemon()
}

func (t *Tea) SteepTeaBag() string {
	return "tea bag steeped"
}

func (t *Tea) AddLemon() string {
	return "lemon added"
}
