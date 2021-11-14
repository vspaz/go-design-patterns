package prototype

type inode interface {
	GetInfo(string string) string
	clone() inode
}
