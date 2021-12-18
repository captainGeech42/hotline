package backend

import (
	"github.com/gorilla/mux"
)

type response struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
}

func ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/callback", newCallback).Methods("POST")
}
