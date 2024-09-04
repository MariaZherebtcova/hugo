package main //мб хватит покрытия?

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReverseProxy(t *testing.T) {
	rp := NewReverseProxy("hugo", "1313")

	req, err := http.NewRequest("GET", "/api/hello", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rr := httptest.NewRecorder()
	handler := rp.ReverseProxy(http.HandlerFunc(handleHelloAPI))
	handler.ServeHTTP(rr, req)

	expected := content
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
	req, err = http.NewRequest("GET", "/static/file.css", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	rr = httptest.NewRecorder()
	handler = rp.ReverseProxy(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("this handler should not be called")
	}))
	handler.ServeHTTP(rr, req)

}
