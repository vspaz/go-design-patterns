package composite

type File struct {
	name string
}

func (f *File) Search(word string) bool {
	return f.name == word
}

func (f *File) GetName() string {
	return f.name
}
