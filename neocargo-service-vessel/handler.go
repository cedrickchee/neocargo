// gRPC handler

package main

import (
	"context"
	"log"

	pb "github.com/cedrickchee/neocargo/neocargo-service-vessel/proto/vessel"
)

// Handler should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type handler struct {
	repository
}

// FindAvailable vessels
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {

	// Find the next available vessel
	vessel, err := s.repository.FindAvailable(ctx, MarshalSpecification(req))
	log.Println("FindAvailable vessels DB call OK")

	if err != nil {
		log.Printf("FindAvailable vessels err: %v", err)
		return err
	}

	// Set the vessel as part of the response message type
	log.Printf("Found vessel: %v", vessel)
	res.Vessel = UnmarshalVessel(vessel)
	return nil
}

// Create a new vessel
func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err := s.repository.Create(ctx, MarshalVessel(req)); err != nil {
		return err
	}
	res.Vessel = req
	return nil
}
