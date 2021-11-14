package prototype

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func generateTestData() (*Dir, *Dir) {
	dir_1 := &Dir{
		children: []Inode{&File{
			name: "file_1",
		}},
		name: "dir_1",
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
		name: "dir_2",
	}
	return dir_1, dir_2
}

func TestGetInfoOk(t *testing.T) {
	dir_1 := &Dir{
		children: []Inode{&File{
			name: "file_1",
		}},
		name: "dir_1",
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
		name: "dir_2",
	}
	assert.Equal(
		t,
		"\ndir_2\ndir_1\nfile_1\nfile_2\nfile_3",
		dir_2.GetInfo("\n"),
	)
}

func TestCloneOk(t *testing.T) {
	dir_1 := &Dir{
		children: []Inode{&File{
			name: "file_1",
		}},
		name: "dir_1",
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
		name: "dir_2",
	}
	assert.Equal(
		t,
		"\ndir_2_clone\ndir_1_clone\nfile_1_clone\nfile_2_clone\nfile_3_clone",
		dir_2.Clone().GetInfo("\n"),
	)
}
