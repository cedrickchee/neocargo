// Interact with MongoDB database

package main

import (
	"context"
	"log"

	pb "github.com/cedrickchee/neocargo/neocargo-service-shipment/proto/shipment"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Datastore models

// Shipment model
type Shipment struct {
	ID          string     `json:"id"`
	Weight      int32      `json:"weight"`
	Description string     `json:"description"`
	Containers  Containers `json:"containers"`
	VesselID    string     `json:"vessel_id"`
}

// Container model
type Container struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	UserID     string `json:"user_id"`
}

// Containers is a list of shipping container
type Containers []*Container

// Marshalling and unmarshalling functions
//
// Convert between the protobuf definition generated structs, and our internal
// datastore models.
//
// You can in theory use the generated structs as your models also, but this
// isn't neccessarily recommended from a software design perspective. As you are
// now coupling your data model, to your delivery layer. It's good to maintain
// these boundaries between the various responsibilities in your software.
// It may seem like additional overhead, but this is important for the
// extensibility of your software.

// MarshalContainerCollection ...
func MarshalContainerCollection(containers []*pb.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}

// UnmarshalContainerCollection ...
func UnmarshalContainerCollection(containers []*Container) []*pb.Container {
	collection := make([]*pb.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}
	return collection
}

// UnmarshalShipmentCollection ...
func UnmarshalShipmentCollection(shipments []*Shipment) []*pb.Shipment {
	collection := make([]*pb.Shipment, 0)
	for _, shipment := range shipments {
		collection = append(collection, UnmarshalShipment(shipment))
	}
	return collection
}

// UnmarshalContainer ...
func UnmarshalContainer(container *Container) *pb.Container {
	return &pb.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

// MarshalContainer ...
func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

// MarshalShipment marshals an input shipment type to a shipment model
func MarshalShipment(shipment *pb.Shipment) *Shipment {
	containers := MarshalContainerCollection(shipment.Containers)
	return &Shipment{
		ID:          shipment.Id,
		Weight:      shipment.Weight,
		Description: shipment.Description,
		Containers:  containers,
		VesselID:    shipment.VesselId,
	}
}

// UnmarshalShipment ...
func UnmarshalShipment(shipment *Shipment) *pb.Shipment {
	return &pb.Shipment{
		Id:          shipment.ID,
		Weight:      shipment.Weight,
		Description: shipment.Description,
		Containers:  UnmarshalContainerCollection(shipment.Containers),
		VesselId:    shipment.VesselID,
	}
}

type repository interface {
	Create(ctx context.Context, shipment *Shipment) error
	GetAll(ctx context.Context) ([]*Shipment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create creates a shipment collection
func (repository *MongoRepository) Create(ctx context.Context, shipment *Shipment) error {
	_, err := repository.collection.InsertOne(ctx, shipment)
	return err
}

// GetAll gets all shipment collection
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Shipment, error) {
	log.Println("Shipment repo.GetAll")

	cur, err := repository.collection.Find(ctx, bson.D{}, nil)
	if err != nil {
		log.Printf("Shipment repo.collection.Find err: %v\n", err)
		return nil, err
	}
	log.Printf("Shipment repo.collection.Find OK. cur: %v", cur)

	var shipments []*Shipment
	for cur.Next(ctx) {
		var shipment *Shipment
		log.Println("Shipment repo cursor")

		if err := cur.Decode(&shipment); err != nil {
			log.Printf("Shipment repo Decode err: %v\n", err)
			return nil, err
		}
		shipments = append(shipments, shipment)
	}

	log.Printf("Shipment repo shipments: %v\n", shipments)
	return shipments, err
}
