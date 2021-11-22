package decorator

type LoggedRequest struct {
	Request Request
}

func (l *LoggedRequest) DoRequest() string {
	return "received request" + l.DoRequest()
}
