package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	tripgrpc "github.com/shemanto27/GoRideSharing/backend/services/api-gateway/internal/grpc"

	tripproto "github.com/shemanto27/GoRideSharing/backend/proto/trip"
)

var tripClient = tripgrpc.NewTripClient()

func GetTripsHandler(w http.ResponseWriter, r *http.Request) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	response, err := tripClient.Client.GetTrip(
		ctx,
		&tripproto.GetTripRequest{
			TripId: "trip-123",
		},
	)

	if err != nil {
		http.Error(
			w,
			"Failed to fetch trip",
			http.StatusInternalServerError,
		)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response)
}