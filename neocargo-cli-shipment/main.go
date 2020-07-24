package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"context"

	pb "github.com/haxorbit/neocargo/neocargo-service-shipment/proto/shipment"

	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
)

const (
	address         = "localhost:50051"
	defaultFilename = "shipment.json"
)

func parseFile(file string) (*pb.Shipment, error) {
	var shipment *pb.Shipment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &shipment)
	return shipment, err
}

func main() {
	// Set up a connection to the server.
	service := micro.NewService(micro.Name("neocargo.cli.shipment"))
	service.Init()

	client := pb.NewShippingService("neocargo.service.shipment", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	var token string
	log.Println(os.Args)

	if len(os.Args) < 3 {
		log.Fatal(errors.New("Not enough arguments, expecting file and token"))
	}

	file = os.Args[1]
	token = os.Args[2]

	shipment, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	// Create a new context which contains our given token.
	// This same context will be passed into both the calls we make
	// to our shipment-service.
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"token": token,
	})

	// First call using our tokenized context
	r, err := client.CreateShipment(ctx, shipment)
	if err != nil {
		log.Fatalf("Could not create a shipment: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	// Second call
	getAll, err := client.GetShipments(ctx, &pb.GetRequest{})
	if err != nil {
		log.Fatalf("Could not list shipments: %v", err)
	}
	for i, v := range getAll.Shipments {
		log.Printf("Shipment %d: %v", i+1, v)
	}
}
