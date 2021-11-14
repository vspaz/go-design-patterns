package prototype

type File struct {
	name string
}

func (f *File) GetInfo(sep string) string {
	return sep + f.name
}
