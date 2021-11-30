package chor

import (
	"go-design-patterns/designpatterns/behavioral/chor/middleware"
	"testing"
	"time"
)

func TestMiddlewareOk(t *testing.T) {
	req := &middleware.Request{
		IsAuthenticated: false,
		User:            "john.doe@example.com",
	}

	loggingMwe := &middleware.Logging{}
	authMwe := &middleware.Auth{
		AllowedUsers: map[string]bool{"john.doe@example.com": true},
	}

	throttlignMwe := &middleware.Throttling{
		CurrentTime: time.Now(),
	}
	throttlignMwe.SetNext(authMwe)
	loggingMwe.SetNext(throttlignMwe)
	loggingMwe.Process(req)
}
