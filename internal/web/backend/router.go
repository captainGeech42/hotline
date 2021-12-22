package backend

import (
	"github.com/captainGeech42/hotline/internal/config"
	"github.com/gorilla/mux"
)

// top level callback domain to be used
var callbackDomain string

func ConfigureRouter(router *mux.Router, cfg *config.Config) {
	router.HandleFunc("/callback", newCallback).Methods("POST")

	callbackDomain = cfg.Server.Callback.Domain
}
