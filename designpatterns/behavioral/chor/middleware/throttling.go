package middleware

import "time"

type ThrottlingMwe struct {
	requestCount int
	rpsLimit     int
	currentTime  time.Time
	next         Middleware
}

func (t *ThrottlingMwe) Process(r *Request) string {
	if time.Now().Second() > t.currentTime.Second()+60 {
		t.requestCount = 0
		t.currentTime = time.Now()
	}
	t.requestCount++

	if t.requestCount < t.rpsLimit {
		return "rps limit exceeded"
	}
	return t.next.Process(r)
}

func (t *ThrottlingMwe) SetNext(next Middleware) {
	t.next = next
}
