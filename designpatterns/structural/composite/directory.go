package composite

type Dir struct {
	name       string
	components []Component
}

func (d *Dir) Search(word string) bool {
	for _, composite := range d.components {
		isMatched := composite.Search(word)
		if isMatched {
			return true
		}
	}
	return false
}

func (d *Dir) Add(component Component) {
	d.components = append(d.components, component)
}
