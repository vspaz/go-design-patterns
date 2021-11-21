package proxy

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const (
	statusEndpoint = "/api/v1/status"
	userEndpoint   = "/api/v1/user"
)

func TestProxyServer(t *testing.T) {
	proxyServer := NewProxyServer(1)
	statusCode, response := proxyServer.handleRequest(statusEndpoint, "GET")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Ok", response)

	statusCode, response = proxyServer.handleRequest(statusEndpoint, "GET")
	assert.Equal(t, http.StatusTooManyRequests, statusCode)
	assert.Equal(t, "Too many requests", response)
}
