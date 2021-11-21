package proxy

type ProxyServer struct {
	AppServer    *WebApp
	RequestCount int
	RateLimiter  map[string]int
}
