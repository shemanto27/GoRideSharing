package main

import (
	"context"
	"log"
	"net"

	tripproto "github.com/shemanto27/GoRideSharing/backend/proto/trip"

	"google.golang.org/grpc"
)

type TripServer struct {
	tripproto.UnimplementedTripServiceServer
}

// RPC Method Implementation, implements generated interface
func (s *TripServer) GetTrip(
	ctx context.Context,
	req *tripproto.GetTripRequest,
) (*tripproto.GetTripResponse, error) {

	log.Printf("Received trip request for trip_id=%s", req.TripId)

	return &tripproto.GetTripResponse{
		TripId: req.TripId,
		Status: "DRIVER_ASSIGNED",
	}, nil
}

func main() {

	// opens TCP socket, gRPC internally uses TCP
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// gRPC runtime server
	grpcServer := grpc.NewServer()

	tripproto.RegisterTripServiceServer(
		grpcServer,
		&TripServer{},
	)

	log.Println("Trip gRPC Service running on :50051")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}