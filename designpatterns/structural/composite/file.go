package composite

import "fmt"

type File struct {
	name string
}

func (f *File) search(word string) string {
	return fmt.Sprintf("searching for word '%s' in file '%s'", word, f.name)
}

func (f *File) getName() string {
	return f.name
}
