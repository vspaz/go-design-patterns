package middleware

type Request struct {
	User            string
	IsAuthenticated bool
}

type Middleware interface {
	Process(r *Request) string
	SetNext(Middleware)
}
