package backend

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(router *mux.Router) {
	router.HandleFunc("/callback", newCallback).Methods("POST")
}
