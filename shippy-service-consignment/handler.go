// gRPC handler

package main

import (
	"context"
	"log"

	pb "github.com/haxorbit/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/haxorbit/shippy/shippy-service-vessel/proto/vessel"
	"github.com/pkg/errors"
)

// Handler should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type handler struct {
	repository
	vesselClient vesselProto.VesselService
}

// CreateConsignment is a method on our service. It creates a new consignment by
// taking a context and a request as an argument. These are handled by the gRPC
// server.
func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error {

	// Here we call a client instance of our vessel service with our consignment
	// weight, and the amount of containers as the capacity value.
	vesselResponse, err := s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity:  int32(len(req.Containers)),
	})
	if vesselResponse == nil {
		return errors.New("error fetching vessel, returned nil")
	}

	if err != nil {
		return err
	}

	// We set the VesselId as the vessel we got back from our vessel service.
	req.VesselId = vesselResponse.Vessel.Id

	// Save our consignment
	if err = s.repository.Create(ctx, MarshalConsignment(req)); err != nil {
		return err
	}

	// Return matching the `Response` message we created in our protobuf
	// definition.
	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments is a method on our service. It gets all consignment by
// taking a context and a request as an argument. These are handled by the gRPC
// server.
func (s *handler) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	log.Println("GetConsignments handler")

	consignments, err := s.repository.GetAll(ctx)
	log.Println("GetConsignments DB call OK")

	if err != nil {
		log.Printf("GetConsignments err: %v", err)
		return err
	}

	log.Printf("Found consignments: %v\n", consignments)
	res.Consignments = UnmarshalConsignmentCollection(consignments)

	log.Println("Unmarshall consignments success")
	return nil
}
