package app

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

		log.Printf("request from %s to %s\n", r.RemoteAddr, r.RequestURI)

		next.ServeHTTP(w, r)
	})
}
