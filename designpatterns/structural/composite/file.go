package composite

type File struct {
	name string
}

func (f *File) Search(word string) string {
	if word == f.name {
		return "found"
	}
	return "not found"
}

func (f *File) GetName() string {
	return f.name
}
