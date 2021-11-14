package prototype

import (
	"fmt"
	"strings"
)

type Dir struct {
	children []Inode
	name     string
}

func (d *Dir) GetInfo(sep string) string {
	var fileNames []string
	fileNames = append(fileNames, sep + d.name)
	for _, file := range d.children {
		fileNames = append(fileNames, file.GetInfo(sep))
	}
	fmt.Println(fileNames)
	return strings.Join(fileNames, "")
}

func (d *Dir) Clone() Inode {
	clonedDir := &Dir{name: d.name + "_clone"}
	var tempChildren []Inode
	for _, file := range d.children {
		clonedFile := file.Clone()
		tempChildren = append(tempChildren, clonedFile)
	}
	clonedDir.children = tempChildren
	return clonedDir
}
