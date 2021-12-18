package app

import (
	"fmt"
	"net/http"

	"github.com/captainGeech42/hotline/internal/app/backend"
	"github.com/captainGeech42/hotline/internal/app/frontend"
	"github.com/captainGeech42/hotline/internal/config"
)

func StartApp(config *config.Config) {
	http.Handle("/api", backend.CreateRouter())
	http.Handle("/", frontend.CreateRouter())

	addr := fmt.Sprintf("0.0.0.0:%d", config.Server.App.Port)

	http.ListenAndServe(addr, nil)
}
