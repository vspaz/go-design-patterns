package decorator

type Auth struct {
	Request Request
}

func (a *Auth) DoRequest() string {
	return "request authenticated; " + a.DoRequest()
}
