package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestDecoratorOk(t *testing.T) {
	newRequest := &Request{}
	assert.Equal(t, "do request; ", newRequest.DoRequest())

	loggedRequest := &LoggedRequest{
		Request: *newRequest,
	}
	assert.Equal(t, "2006-02-08 22:20:02,165: do request; ", loggedRequest.DoRequest())
}