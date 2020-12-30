package http

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/test/123456", req.URL.String())
		// Send response to be tested
		_, _ = rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	req := NewClient()
	body, err := req.Get(server.URL + "/test/123456")

	assert.NoError(t, err)
	assert.Equal(t, "OK", body)
}

func TestClient_Post(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, "POST", req.Method)
		body, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, "{\"property\":\"test\"}", string(body))
		assert.Equal(t, "/test", req.URL.String())
		// Send response to be tested
		_, _ = rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	req := NewClient()
	mockReq := &PostMock{
		Property: "test",
	}
	body, err := req.Post(server.URL+"/test", &mockReq)

	assert.NoError(t, err)
	assert.Equal(t, "OK", body)
}

func TestClient_Delete(t *testing.T) {
	// Start a local HTTP server
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// Test request parameters
		assert.Equal(t, "DELETE", req.Method)
		assert.Equal(t, "/test/123456", req.URL.String())
		// Send response to be tested
		_, _ = rw.Write([]byte(`OK`))
	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	req := NewClient()
	err := req.Delete(server.URL + "/test/123456")
	assert.NoError(t, err)
}

type PostMock struct {
	Property string `json:"property"`
}
