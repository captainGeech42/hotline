package web

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/captainGeech42/hotline/internal/config"
	"github.com/captainGeech42/hotline/internal/db"
	"github.com/captainGeech42/hotline/internal/web/backend"
	"github.com/captainGeech42/hotline/internal/web/frontend"
	"github.com/gorilla/mux"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("app good to go"))
}

// Start the webapp HTTP server. This handled the frontend and backend API routes
func StartApp(cfg *config.Config) {
	// seed rand for callback name generation
	rand.Seed(time.Now().UnixNano())

	// define the top level router
	router := mux.NewRouter()
	router.Use(loggingMiddleware)
	router.HandleFunc("/healthcheck", healthCheck)

	// create subrouters for the backend and frontend packages
	backendRouter := router.PathPrefix("/api").Subrouter()
	frontendRouter := router.PathPrefix("/").Subrouter()

	// configure their specific routes/middlware
	backend.ConfigureRouter(backendRouter)
	frontend.ConfigureRouter(frontendRouter)

	// connect to database
	if !db.ConnectToDb(cfg.Server.Database) {
		return
	}

	// start the http listener
	addr := fmt.Sprintf("0.0.0.0:%d", cfg.Server.Web.Port)
	log.Printf("listening on %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, router))
}
