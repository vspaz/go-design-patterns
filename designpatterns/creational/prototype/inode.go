package prototype

type inode interface {
	GetInfo(string string)
	clone() inode
}
