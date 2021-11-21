package proxy

import "net/http"

type ProxyServer struct {
	AppServer    *WebApp
	RequestCount int
	RateLimiter  map[string]int
}

func NewProxyServer(requestLimit int) *ProxyServer {
	return &ProxyServer{
		AppServer:    &WebApp{},
		RequestCount: requestLimit,
		RateLimiter:  make(map[string]int),
	}
}

func (p *ProxyServer) isRequestLimitHit(endpoint string) bool {
	p.RateLimiter[endpoint]++
	return p.RateLimiter[endpoint] > p.RequestCount
}

func (p *ProxyServer) handleRequest(endpoint string, method string) (int, string) {
	if p.isRequestLimitHit(endpoint) {
		return http.StatusTooManyRequests, "Too many requests"
	}
	return p.AppServer.handleRequest(endpoint, method)
}
