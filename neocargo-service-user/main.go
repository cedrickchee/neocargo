package main

import (
	"log"

	pb "github.com/haxorbit/neocargo/neocargo-service-user/proto/user"
	"github.com/micro/go-micro/v2"
)

// SQL DDL for creating new schema.
const schema = `
	create table if not exists users (
		id varchar(36) not null,
		name varchar(125) not null,
		email varchar(225) not null unique,
		password varchar(225) not null,
		company varchar(125),
		primary key (id)
	);
`

func main() {

	// Creates a database connection and handles closing it again before exit.
	db, err := NewConnection()
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	if err != nil {
		log.Fatalf("Could not connect to PostgreSQL database: %v", err)
	}

	// Initialize database

	// Run schema query on start-up, as we're using "create if not exists"
	// this will only be ran once. In order to create updates, you'll need to
	// use a migrations library.
	db.MustExec(schema)

	repo := NewPostgresRepository(db)

	tokenService := &TokenService{repo}

	// Set up micro instance
	service := micro.NewService(
		micro.Name("neocargo.service.user"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	service.Init()

	h := &handler{repo, tokenService}

	// Register handlers
	if err := pb.RegisterUserServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	// Run the server
	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
