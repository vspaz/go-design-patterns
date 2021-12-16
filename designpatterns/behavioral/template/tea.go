package template

type Tea struct {
}

func (t *Tea) SteepTeaBag() string {
	return "tea bag steeped"
}

func (t *Tea) AddLemon() string {
	return "lemon added"
}
