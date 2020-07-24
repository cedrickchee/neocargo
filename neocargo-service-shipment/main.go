package main

import (
	"context"
	"log"
	"os"

	// Import the generated protobuf code
	pb "github.com/haxorbit/neocargo/neocargo-service-shipment/proto/shipment"
	userProto "github.com/haxorbit/neocargo/neocargo-service-user/proto/user"
	vesselProto "github.com/haxorbit/neocargo/neocargo-service-vessel/proto/vessel"
	"github.com/pkg/errors"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
)

const (
	defaultHost = "datastore:27017"
)

var (
	service micro.Service
)

// AuthWrapper is a handler wrapper - a high-order function which takes a HandlerFunc
// and returns a function, which takes a context, request and response interface.
// The token is extracted from the context set in our shipment-cli, that
// token is then sent over to the user service to be validated.
// If valid, the call is passed along to the handler. If not,
// an error is returned.
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}

		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("No auth meta-data found in request")
		}

		// Note this is now uppercase (not entirely sure why this is...)
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Auth here
		// Really shouldn't be using a global here, find a better way of doing
		// this, since you can't pass it into a wrapper.
		authClient := userProto.NewUserService("neocargo.service.user", service.Client())
		_, err := authClient.ValidateToken(context.Background(), &userProto.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}

func main() {

	// Set up micro instance
	service = micro.NewService(

		// This name must match the package name given in your protobuf definition
		micro.Name("neocargo.service.shipment"),
		// Our auth middleware
		micro.WrapHandler(AuthWrapper),
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
	consignmentCollection := client.Database("neocargo").Collection("consignments")
	repository := &MongoRepository{consignmentCollection}

	vesselClient := vesselProto.NewVesselService("neocargo.service.vessel", service.Client())

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
