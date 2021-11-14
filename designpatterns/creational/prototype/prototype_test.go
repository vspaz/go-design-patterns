package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetInfoOk(t *testing.T) {
	dir_1 := &Dir{
		children: []Inode{&File{
			name: "dir_1",
		}},
	}
	dir_2 := &Dir{
		children: []Inode{
			dir_1,
			&File{
				name: "file_2",
			},
			&File{
				name: "file_3",
			},
		},
	}
	assert.Equal(t, "\n\n\ndir_1\nfile_2\nfile_3", dir_2.GetInfo("\n"))
}
