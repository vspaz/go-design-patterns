package middleware

import (
	"log"
	"time"
)

type Logging struct {
	next Middleware
}

func (l *Logging) Process(r *Request) {
	log.Printf(" %s: request received", time.Now().Format("2006-01-02 15:04:05.000"))
	l.next.Process(r)
}

func (l *Logging) setNext(next Middleware) {
	l.next = next
}
