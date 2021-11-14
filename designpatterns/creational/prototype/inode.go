package prototype

type Inode interface {
	GetInfo(string string) string
	Clone() Inode
}
