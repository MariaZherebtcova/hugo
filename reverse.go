package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

type ReverseProxy struct {
	target *url.URL
	proxy  *httputil.ReverseProxy
}

func NewReverseProxy(host, port string) *ReverseProxy {
	targetURL := &url.URL{
		Scheme: "http",
		Host:   host + ":" + port,
	}
	return &ReverseProxy{
		target: targetURL,
		proxy:  httputil.NewSingleHostReverseProxy(targetURL),
	}
}

func (rp *ReverseProxy) ReverseProxy(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			next.ServeHTTP(w, r)
			return
		} else {
			rp.proxy.ServeHTTP(w, r)
		}

	})
}
