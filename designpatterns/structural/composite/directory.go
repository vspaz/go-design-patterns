package composite

type Dir struct {
	name string
	components []Component
}

func (d *Dir) Search(word string) string {
	for _, composite := range d.components {
		found := composite.Search(word)
		if found == "found" {
			return found
		}
	}
	return "not found"
}

func (d *Dir) Add(component Component) {
	d.components = append(d.components, component)
}
