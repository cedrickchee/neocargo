// Interact with MongoDB database

package main

import (
	"context"
	"log"

	pb "github.com/haxorbit/shippy/shippy-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

// Datastore models

// Consignment model
type Consignment struct {
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

// UnmarshalConsignmentCollection ...
func UnmarshalConsignmentCollection(consignments []*Consignment) []*pb.Consignment {
	collection := make([]*pb.Consignment, 0)
	for _, consignment := range consignments {
		collection = append(collection, UnmarshalConsignment(consignment))
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

// MarshalConsignment marshals an input consignment type to a consignment model
func MarshalConsignment(consignment *pb.Consignment) *Consignment {
	containers := MarshalContainerCollection(consignment.Containers)
	return &Consignment{
		ID:          consignment.Id,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  containers,
		VesselID:    consignment.VesselId,
	}
}

// UnmarshalConsignment ...
func UnmarshalConsignment(consignment *Consignment) *pb.Consignment {
	return &pb.Consignment{
		Id:          consignment.ID,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  UnmarshalContainerCollection(consignment.Containers),
		VesselId:    consignment.VesselID,
	}
}

type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

// MongoRepository implementation
type MongoRepository struct {
	collection *mongo.Collection
}

// Create creates a consignment collection
func (repository *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := repository.collection.InsertOne(ctx, consignment)
	return err
}

// GetAll gets all consignment collection
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error) {
	log.Println("Consignment repo.GetAll")

	cur, err := repository.collection.Find(ctx, nil, nil)
	if err != nil {
		log.Printf("Consignment repo.collection.Find err: %v\n", err)
		return nil, err
	}
	log.Printf("Consignment repo.collection.Find OK. cur: %v", cur)

	var consignments []*Consignment
	for cur.Next(ctx) {
		var consignment *Consignment
		log.Println("Consignment repo cursor")

		if err := cur.Decode(&consignment); err != nil {
			log.Printf("Consignment repo Decode err: %v\n", err)
			return nil, err
		}
		consignments = append(consignments, consignment)
	}

	log.Printf("Consignment repo consignments: %v\n", consignments)
	return consignments, err
}
