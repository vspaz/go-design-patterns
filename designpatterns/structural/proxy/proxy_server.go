package proxy

type ProxyServer struct {
	AppServer    *WebApp
	RequestCount int
	RateLimiter  map[string]int
}

func NewProxyServer() *ProxyServer {
	return &ProxyServer{
		AppServer:    &WebApp{},
		RequestCount: 5,
		RateLimiter:  make(map[string]int),
	}
}
