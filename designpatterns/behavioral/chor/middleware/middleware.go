package middleware

type Request struct {}

type Middleware interface {
	Process(*Request)
	setNext(Middleware)
}