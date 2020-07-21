package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/haxorbit/shippy/shippy-service-user/proto/user"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
)

func createUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("shippy.service.user", service.Client())

	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	// Print the response
	fmt.Println("Created user. Response: ", rsp.User)

	return nil
}

func main() {
	// Create and initialize a new service
	service := micro.NewService()

	service.Init(
		micro.Action(func(c *cli.Context) error {

			// Using go-micro's command line helper
			name := c.String("name")
			email := c.String("email")
			company := c.String("company")
			password := c.String("password")

			ctx := context.Background()
			user := &pb.User{
				Name:     name,
				Email:    email,
				Company:  company,
				Password: password,
			}

			if err := createUser(ctx, service, user); err != nil {
				log.Println("Error creating user: ", err.Error())
				return err
			}

			return nil
		}),
	)
}
