package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/captainGeech42/hotline/internal/app/backend"
	"github.com/captainGeech42/hotline/internal/app/frontend"
	"github.com/captainGeech42/hotline/internal/config"
	"github.com/gorilla/mux"
)

// Start the webapp HTTP server. This handled the frontend and backend API routes
func StartApp(config *config.Config) {
	// define the top level router
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	// create subrouters for the backend and frontend packages
	backendRouter := router.PathPrefix("/api").Subrouter()
	frontendRouter := router.PathPrefix("/").Subrouter()

	// configure their specific routes/middlware
	backend.ConfigureRouter(backendRouter)
	frontend.ConfigureRouter(frontendRouter)

	// start the http listener
	addr := fmt.Sprintf("0.0.0.0:%d", config.Server.App.Port)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
