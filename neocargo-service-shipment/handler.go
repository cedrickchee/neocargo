// gRPC handler

package main

import (
	"context"
	"log"

	pb "github.com/cedrickchee/neocargo/neocargo-service-shipment/proto/shipment"
	vesselProto "github.com/cedrickchee/neocargo/neocargo-service-vessel/proto/vessel"
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

// CreateShipment is a method on our service. It creates a new shipment by
// taking a context and a request as an argument. These are handled by the gRPC
// server.
func (s *handler) CreateShipment(ctx context.Context, req *pb.Shipment, res *pb.Response) error {

	// Here we call a client instance of our vessel service with our shipment
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

	// Save our shipment
	if err = s.repository.Create(ctx, MarshalShipment(req)); err != nil {
		return err
	}

	// Return matching the `Response` message we created in our protobuf
	// definition.
	res.Created = true
	res.Shipment = req
	return nil
}

// GetShipments is a method on our service. It gets all shipment by
// taking a context and a request as an argument. These are handled by the gRPC
// server.
func (s *handler) GetShipments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	log.Println("GetShipments handler")

	shipments, err := s.repository.GetAll(ctx)
	log.Println("GetShipments DB call OK")

	if err != nil {
		log.Printf("GetShipments err: %v", err)
		return err
	}

	log.Printf("Found shipments: %v\n", shipments)
	res.Shipments = UnmarshalShipmentCollection(shipments)

	log.Println("Unmarshall shipments success")
	return nil
}
