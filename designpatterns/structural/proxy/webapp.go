package proxy

import "net/http"

type WebApp struct{}

func (w *WebApp) handleRequest(endpoint string, method string) (int, string) {
	if endpoint == "/api/v1/status" && method == http.MethodHead {
		return http.StatusOK, "Ok"
	} else if endpoint == "/api/v1/user" && method == http.MethodPost {
		return http.StatusAccepted, "Accepted"
	}
	return http.StatusBadRequest, "Bad request"
}
