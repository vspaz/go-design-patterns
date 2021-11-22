package decorator

type ProfileRequest struct {
	Request IRequest
}

func (p *ProfileRequest) DoRequest() string {
	return "request took: '0.02ms to process'; " + p.Request.DoRequest()
}
