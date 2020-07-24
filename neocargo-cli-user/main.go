package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/haxorbit/neocargo/neocargo-service-user/proto/user"

	"github.com/micro/cli/v2"
	micro "github.com/micro/go-micro/v2"
)

func createUser(ctx context.Context, service micro.Service, user *pb.User) error {
	client := pb.NewUserService("neocargo.service.user", service.Client())

	rsp, err := client.Create(ctx, user)
	if err != nil {
		return err
	}

	// Print the response
	fmt.Println("Created user. Response: ", rsp.User)

	return nil
}

func listUsers(ctx context.Context, service micro.Service) error {
	client := pb.NewUserService("neocargo.service.user", service.Client())

	getAll, err := client.GetAll(ctx, &pb.Request{})
	if err != nil {
		return err
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	return nil
}

func authenticate(ctx context.Context, service micro.Service, email string, password string) error {
	client := pb.NewUserService("neocargo.service.user", service.Client())

	authResponse, err := client.Auth(ctx, &pb.User{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return err
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)
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

			if err := listUsers(ctx, service); err != nil {
				log.Println("Could not list users: ", err.Error())
				return err
			}

			if err := authenticate(context.TODO(), service, email, password); err != nil {
				log.Printf("Could not authenticate user: %s error: %v\n", email, err.Error())
				return err
			}

			return nil
		}),
	)
}
