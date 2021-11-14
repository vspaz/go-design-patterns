package prototype

type File struct {
	name string
}

func (f *File) GetInfo(sep string) string {
	return sep + f.name
}

func (f *File) Clone() Inode {
	return &File{name: f.name + "_clone"}
}
