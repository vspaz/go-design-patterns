package decorator

type Auth struct {
	Request IRequest
}

func (a *Auth) DoRequest() string {
	return "request authenticated; " + a.Request.DoRequest()
}
