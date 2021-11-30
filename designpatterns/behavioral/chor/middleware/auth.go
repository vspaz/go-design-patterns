package middleware

type Auth struct {
	next         Middleware
	AllowedUsers map[string]bool
}

func (a *Auth) Process(r *Request) {
	if _, hasUser := a.AllowedUsers[r.User]; hasUser {
		a.next.Process(r)
	}
	r.IsAuthenticated = false
	a.next.Process(r)
}

func (a *Auth) SetNext(next Middleware) {
	a.next = next
}
