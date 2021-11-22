package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestDecoratorOk(t *testing.T) {
	newRequest := &Request{}
	assert.Equal(t, "do request; ", newRequest.DoRequest(), )

}