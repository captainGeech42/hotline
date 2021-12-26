package frontend

import (
	"net/http"

	"github.com/gorilla/mux"
)

func ErrorResp(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Write([]byte("If you're seeing this, your NGINX routing isn't setup correctly. Please double check your config to make sure you're serving the React app properly"))
}

func ConfigureRouter(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(ErrorResp)
}
