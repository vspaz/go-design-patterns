package middleware

type Request struct {
	user            string
	IsAuthenticated bool
}

type Middleware interface {
	Process(*Request) string
	SetNext(Middleware)
}
