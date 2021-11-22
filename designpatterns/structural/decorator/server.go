package decorator

type IRequest interface {
	DoRequest() string
}

type Request struct{}

func (r *Request) DoRequest() string {
	return "do request; "
}
