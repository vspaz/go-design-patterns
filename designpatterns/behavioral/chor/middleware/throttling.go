package middleware

import (
	"log"
	"time"
)

type Throttling struct {
	requestCount int
	rpsLimit     int
	CurrentTime  time.Time
	next         Middleware
}

func (t *Throttling) Process(r *Request) {
	if time.Now().Second() > t.CurrentTime.Second()+60 {
		t.requestCount = 0
		t.CurrentTime = time.Now()
	}
	t.requestCount++

	if t.requestCount < t.rpsLimit {
		log.Println("rps limit exceeded")
	}
	t.next.Process(r)
}

func (t *Throttling) SetNext(next Middleware) {
	t.next = next
}
