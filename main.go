package main

import (
	"github.com/go-chi/chi"
	"net/http"
)

const content = `Hello from API`

func main() {
	r := chiRouter()
	hugo := NewReverseProxy("hugo", "1313")
	r.Use(hugo.ReverseProxy)
	r.Get("/api/*", handleHelloAPI)
	http.ListenAndServe(":8080", r)
}

func chiRouter() *chi.Mux {
	r := chi.NewRouter()
	return r
}
func handleHelloAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(content))
}
