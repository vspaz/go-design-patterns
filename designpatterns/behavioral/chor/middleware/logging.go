package middleware

import (
	"log"
	"time"
)

type Logging struct {
	next Middleware
}

func (l *Logging) Process(r *Request) string {
	log.Printf("%s: request received\n", time.Now().Format("2006-01-02 15:04:05.000"))
	ret := l.next.Process(r)
	return "request logged; " + ret
}

func (l *Logging) SetNext(next Middleware) {
	l.next = next
}
