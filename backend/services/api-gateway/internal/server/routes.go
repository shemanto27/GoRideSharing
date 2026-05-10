package server

import (
	"net/http"

	"github.com/shemanto27/GoRideSharing/backend/services/api-gateway/internal/handlers"
	"github.com/shemanto27/GoRideSharing/backend/services/api-gateway/internal/middleware"
)

func RegisterRoutes(mux *http.ServeMux) {

	healthHandler := middleware.Chain(
		http.HandlerFunc(handlers.Healthhandler),
		middleware.Logging,
		middleware.CORS,
	)

	tripsHandler := middleware.Chain(
		http.HandlerFunc(handlers.GetTripsHandler),
		middleware.Logging,
		middleware.CORS,
	)

	mux.Handle("/health", healthHandler)

	mux.Handle("/api/v1/trips", tripsHandler)
}