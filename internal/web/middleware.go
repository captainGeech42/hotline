package web

import (
	"log"
	"net/http"
)

// basic logging middleware
// based on the gorilla/mux docs
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// don't log healthcheck b/c i don't care
		if r.RequestURI == "/healthcheck" {
			return
		}

		// nginx setup sets the X-Hotline-Ip header to show the real IP without
		// colliding with a another X-Proxy header. if that is set, we don't
		// need to worry about logging requests
		if _, exists := r.Header["X-Hotline-Real-Ip"]; !exists {
			log.Printf("request from %s to %s\n", r.RemoteAddr, r.RequestURI)
		}

		next.ServeHTTP(w, r)
	})
}
