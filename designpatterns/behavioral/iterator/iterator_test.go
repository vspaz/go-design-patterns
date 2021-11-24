package iterator

import "testing"

func TestIteratorOk(t *testing.T) {
	users := &UserCollection{[]*User{
		{
			Name: "noname",
			Age:  30,
		},
	},
	}
}
