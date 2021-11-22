package decorator

type LoggedRequest struct {
	Request Request
}

func (l *LoggedRequest) DoRequest() string {
	return "2006-02-08 22:20:02,165: " + l.Request.DoRequest()
}
