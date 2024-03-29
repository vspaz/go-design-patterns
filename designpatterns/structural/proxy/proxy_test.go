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

func TestStatusGetHandler(t *testing.T) {
	proxyServer := NewProxyServer(1)
	statusCode, response := proxyServer.handleRequest(statusEndpoint, "GET")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Ok", response)

	statusCode, response = proxyServer.handleRequest(statusEndpoint, "GET")
	assert.Equal(t, http.StatusTooManyRequests, statusCode)
	assert.Equal(t, "Too many requests", response)
}

func TestUserPostHandler(t *testing.T) {
	proxyServer := NewProxyServer(1)
	statusCode, response := proxyServer.handleRequest(userEndpoint, "POST")
	assert.Equal(t, http.StatusAccepted, statusCode)
	assert.Equal(t, "Accepted", response)

	statusCode, response = proxyServer.handleRequest(userEndpoint, "POST")
	assert.Equal(t, http.StatusTooManyRequests, statusCode)
	assert.Equal(t, "Too many requests", response)
}

func TestUnsupportedMethod(t *testing.T) {
	proxyServer := NewProxyServer(1)
	statusCode, response := proxyServer.handleRequest(userEndpoint, "HEAD")
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, "Bad request", response)

	statusCode, response = proxyServer.handleRequest(userEndpoint, "POST")
	assert.Equal(t, http.StatusTooManyRequests, statusCode)
	assert.Equal(t, "Too many requests", response)

}

func TestUndefinedEndpoint(t *testing.T) {
	proxyServer := NewProxyServer(1)
	statusCode, response := proxyServer.handleRequest("someInvalidEndpoint", "GET")
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.Equal(t, "Bad request", response)

	statusCode, response = proxyServer.handleRequest("someInvalidEndpoint", "GET")
	assert.Equal(t, http.StatusTooManyRequests, statusCode)
	assert.Equal(t, "Too many requests", response)
}

func TestDifferentRequests(t *testing.T) {
	proxyServer := NewProxyServer(1)
	statusCode, response := proxyServer.handleRequest(statusEndpoint, "GET")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.Equal(t, "Ok", response)

	statusCode, response = proxyServer.handleRequest(userEndpoint, "POST")
	assert.Equal(t, http.StatusAccepted, statusCode)
	assert.Equal(t, "Accepted", response)
}
