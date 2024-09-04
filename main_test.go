package main

import (
	"github.com/go-chi/chi"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMainFunction(t *testing.T) {
	go func() {
		main()
	}()

	time.Sleep(1 * time.Second) // Даем время серверу запуститься

	r := chi.NewRouter()
	hugo := NewReverseProxy("proxy", "8080")
	r.Use(hugo.ReverseProxy)

	req, _ := http.NewRequest("GET", "/api/", nil)

	testCases := []struct {
		req          *http.Request
		expectedCode int
		expectedBody string
	}{
		{req, http.StatusOK, "Hello from API"},
	}
	for _, tc := range testCases {
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, tc.req)

	}
}
func TestHandleHelloAPI(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handleHelloAPI(rr, req)

	if rr.Body.String() != content {
		t.Errorf("Expected response body '%s', got '%s'", content, rr.Body.String())
	}
}

func TestChiRouter(t *testing.T) {
	r := chiRouter()
	if r == nil {
		t.Errorf("chiRouter() returned nil")
	}
}
