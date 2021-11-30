package middleware

import (
	"time"
)

type Throttling struct {
	requestCount int
	rpsLimit     int
	CurrentTime  time.Time
	next         Middleware
}

func (t *Throttling) Process(r *Request) string {
	if time.Now().Second() > t.CurrentTime.Second()+60 {
		t.requestCount = 0
		t.CurrentTime = time.Now()
	}
	t.requestCount++

	if t.requestCount < t.rpsLimit {
		return "rps limit exceeded; "
	}
	ret := t.next.Process(r)
	return "checked if rps limit exceeded; " + ret
}

func (t *Throttling) SetNext(next Middleware) {
	t.next = next
}
