package composite

type Dir struct {
	components []Component
}

func (d *Dir) Search(word string) string {
	for _, composite := range d.components {
		found := composite.Search(word)
		if found != "" {
			return found
		}
	}
	return ""
}

func (d *Dir) Add(component Component) {
	d.components = append(d.components, component)
}
