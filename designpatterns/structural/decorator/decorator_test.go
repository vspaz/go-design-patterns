package decorator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestDecoratorOk(t *testing.T) {
	newRequest := &Request{}
	assert.Equal(t, "do request; ", newRequest.DoRequest())

	profiledRequest := &ProfileRequest{
		Request: newRequest,
	}
	assert.Equal(t, "request took: '0.02ms to process'; do request; ", profiledRequest.DoRequest())
	authenticatedRequest := &Auth{
		Request: profiledRequest,
	}
	assert.Equal(t, "request authenticated; request took: '0.02ms to process'; do request; ", authenticatedRequest.DoRequest())

	loggedRequest := &LoggedRequest{
		Request: authenticatedRequest,
	}
	assert.Equal(t, "2006-02-08 22:20:02,165: request authenticated; request took: '0.02ms to process'; do request; ", loggedRequest.DoRequest())

}
