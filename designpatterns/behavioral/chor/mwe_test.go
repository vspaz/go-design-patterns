package chor

import (
	"github.com/stretchr/testify/assert"
	"go-design-patterns/designpatterns/behavioral/chor/middleware"
	"testing"
	"time"
)

func TestMiddlewareOk(t *testing.T) {
	authMwe := &middleware.Auth{
		AllowedUsers: map[string]bool{"john.doe@example.com": true},
	}
	throttlingMwe := &middleware.Throttling{
		CurrentTime: time.Now(),
	}
	throttlingMwe.SetNext(authMwe)
	loggingMwe := &middleware.Logging{}
	loggingMwe.SetNext(throttlingMwe)
	assert.Equal(t,
		"request logged; checked if rps limit exceeded; request authenticated",
		loggingMwe.Process(&middleware.Request{
			IsAuthenticated: false,
			User:            "john.doe@example.com",
		}),
	)
}
