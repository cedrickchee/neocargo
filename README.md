# Shippy Microservices in Go

Shippy is a shipping container management platform.

This project is an example of multiple microservices implementation in Go in a monorepo.

The Shippy backend consists of 3 microservices:

- [Consignments](./shippy-service-consignment)
- [Vessels](./shippy-service-vessel)
- [Users](./shippy-service-user)
- Authentication

## Tech Stack

- [Protocol Buffers](https://developers.google.com/protocol-buffers) and [gRPC](https://grpc.io/) as transport protocol
- [go-micro](https://micro.mu/)
- Docker container ([Alpine Linux](https://alpinelinux.org/about/) as base image)
- Docker Compose
- PostgreSQL or MongoDB database
- Google Cloud
- Kubernetes
- NATS
- CircleCI
- Terraform

## Prerequisite

- This project assumes you're using Go 1.13 and upwards and have `export GO111MODULE=on` set.
- Install Go
- Install the protoc compiler
- Install gRPC / protobuf
- Install Go libraries:
    - [protoc-gen-go](https://pkg.go.dev/github.com/golang/protobuf/protoc-gen-go): The protoc-gen-go binary is a protoc plugin to generate a Go protocol buffer package.
