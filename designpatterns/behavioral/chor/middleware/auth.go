package middleware

type Auth struct {
	next         Middleware
	allowedUsers map[string]string
}

func (a *Auth) Process(r *Request) {
	if _, hasUser := a.allowedUsers[r.user]; hasUser {
		a.next.Process(r)
	}
	r.IsAuthenticated = false
	a.next.Process(r)
}

func (a *Auth) SetNext(next Middleware) {
	a.next = next
}
