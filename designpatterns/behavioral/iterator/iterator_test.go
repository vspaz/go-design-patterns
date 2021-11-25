package iterator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIteratorOk(t *testing.T) {
	users := &UserCollection{
		[]*User{
			{
				Name: "noname",
				Age:  30,
			},
			{
				Name: "anonymous",
				Age:  10,
			},
		},
	}

	iterator := users.CreateIterator()
	count := 0
	for iterator.HasNext() {
		user := iterator.Next()
		fmt.Println(user.Name)
		count++
	}
	assert.Equal(t, 2, count)
}
