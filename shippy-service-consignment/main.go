package main

import (
	"context"
	"log"
	"os"

	// Import the generated protobuf code
	pb "github.com/haxorbit/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/haxorbit/shippy/shippy-service-vessel/proto/vessel"

	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	// Set up micro instance
	service := micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.service.consignment"),
	)

	// Init will parse the command line flags.
	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}

	// Create database client
	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	// Initialize database and collection
	consignmentCollection := client.Database("shippy").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}

	vesselClient := vesselProto.NewVesselService("shippy.service.vessel", service.Client())

	h := &handler{repository, vesselClient}

	// Register handlers
	if err := pb.RegisterShippingServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
