package composite

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDirSearchOk(t *testing.T) {
	file_1 := &File{name: "file_1"}
	file_2 := &File{name: "file_2"}
	file_3 := &File{name: "file_3"}

	dir_1 := &Dir{name: "dir_1"}
	dir_1.Add(file_1)

	dir_2 := &Dir{name: "dir_2"}
	dir_2.Add(file_2)
	dir_2.Add(file_3)
	dir_2.Add(dir_1)

	assert.True(t, dir_2.Search("file_1"))
	assert.False(t, dir_2.Search("non existing file"))
}
