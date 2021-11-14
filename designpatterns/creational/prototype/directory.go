package prototype

import "strings"

type Dir struct {
	children []inode
	name     string
}

func (d *Dir) GetInfo(sep string) string {
	fileNames := []string{sep + d.name}
	for _, file := range d.children {
		fileNames = append(fileNames, file.GetInfo(sep+sep))
	}
	return strings.Join(fileNames, "")
}
