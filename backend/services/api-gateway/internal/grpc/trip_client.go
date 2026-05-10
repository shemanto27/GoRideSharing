package grpc

import (
	"log"

	tripproto "github.com/shemanto27/GoRideSharing/backend/proto/trip"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type TripClient struct {
	Client tripproto.TripServiceClient
}

func NewTripClient() *TripClient {

	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatalf("Failed to connect to Trip Service: %v", err)
	}

	client := tripproto.NewTripServiceClient(conn)

	return &TripClient{
		Client: client,
	}
}