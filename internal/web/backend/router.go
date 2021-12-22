package backend

import (
	"github.com/captainGeech42/hotline/internal/config"
	"github.com/gorilla/mux"
)

// top level callback domain to be used
var callbackDomain string

func ConfigureRouter(router *mux.Router, cfg *config.Config) {
	// add routes
	router.HandleFunc("/callback", newCallback).Methods("POST")
	router.HandleFunc("/callback/requests", getCbRequests).Methods("GET")

	// set globals
	callbackDomain = cfg.Server.Callback.Domain
}
