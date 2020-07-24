package main

import (
	"context"
	"log"
	"os"

	pb "github.com/haxorbit/neocargo/neocargo-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

func main() {

	// Set up micro instance
	service := micro.NewService(
		micro.Name("neocargo.service.vessel"),
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
	vesselCollection := client.Database("neocargo").Collection("vessels")
	repository := &MongoRepository{vesselCollection}

	h := &handler{repository}

	// Register handlers
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
