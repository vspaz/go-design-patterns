package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestDecoratorOk(t *testing.T) {
	newRequest := &Request{}
	assert.Equal(t, "do request; ", newRequest.DoRequest())

	authenticatedRequest := &Auth{
		Request: newRequest,
	}
	assert.Equal(t, "request authenticated; do request; ", authenticatedRequest.DoRequest())

	loggedRequest := &LoggedRequest{
		Request: authenticatedRequest,
	}
	assert.Equal(t, "2006-02-08 22:20:02,165: request authenticated; do request; ", loggedRequest.DoRequest())

}
