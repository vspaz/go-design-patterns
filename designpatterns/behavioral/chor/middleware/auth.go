package middleware

type Auth struct {
	next         Middleware
	AllowedUsers map[string]bool
}

func (a *Auth) Process(r *Request) string {
	if _, hasUser := a.AllowedUsers[r.User]; hasUser {
		return "request authenticated"
	}
	r.IsAuthenticated = false
	return "request unauthorized"
}

func (a *Auth) SetNext(next Middleware) {
	a.next = next
}
