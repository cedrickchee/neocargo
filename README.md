# neoCargo Microservices in Go

neoCargo is a shipping container management platform.

This project is an example of multiple microservices implementation in Go in a monorepo.

The neoCargo backend consists of 3 microservices:

- [Shipments](./neocargo-service-shipment)
- [Vessels](./neocargo-service-vessel)
- [Users](./neocargo-service-user)
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

## Usage

### Installation

- Rename file `config.env.example` in project root to `config.env`
- Replace the variables (e.g.: `POSTGRES_PASSWORD`) in `config.env`

### Build and Run Docker Compose Stack

I created a simple Makefile to build, run, test, and teardown Docker Compose
stack in local development machine.

Just run in project root:

```sh
# Build stack
$ make build

# Run stack
$ make run

docker-compose up
Creating network "neocargo_default" with the default driver
Creating database          ... done
Creating datastore ... done
Creating neocargo_vessel_1 ... done
Creating neocargo_user_1     ... done
Creating neocargo_shipment_1 ... done
Creating neocargo_user-cli_1 ... done
Creating neocargo_cli_1      ... done
Attaching to datastore, database, neocargo_vessel_1, neocargo_user_1, neocargo_shipment_1, neocargo_user-cli_1, neocargo_cli_1
database     | The files belonging to this database system will be owned by user "postgres".
database     | This user must also own the server process.
database     |
database     | The database cluster will be initialized with locale "en_US.utf8".
database     | The default database encoding has accordingly been set to "UTF8".
database     | The default text search configuration will be set to "english".
database     |
database     | Data page checksums are disabled.
database     |
database     | fixing permissions on existing directory /var/lib/postgresql/data ... ok
database     | creating subdirectories ... ok
database     | selecting dynamic shared memory implementation ... posix
database     | selecting default max_connections ... 100
database     | selecting default shared_buffers ... 128MB
database     | selecting default time zone ... UTC
database     | creating configuration files ... ok
database     | running bootstrap script ... ok
database     | performing post-bootstrap initialization ... sh: locale: not found
database     | 2020-07-24 14:32:04.281 UTC [29] WARNING:  no usable system locales were found
database     | ok
database     | syncing data to disk ... initdb: warning: enabling "trust" authentication for local connections
database     | You can change this by editing pg_hba.conf or using the option -A, or
database     | --auth-local and --auth-host, the next time you run initdb.
database     | ok
database     |
database     |
database     | Success. You can now start the database server using:
database     |
database     |     pg_ctl -D /var/lib/postgresql/data -l logfile start
database     |
database     | waiting for server to start....2020-07-24 14:32:05.542 UTC [34] LOG:  starting PostgreSQL 12.3 on x86_64-pc-linux-musl, compiled by gcc (Alpine 9.3.0) 9.3.0, 64-bit
database     | 2020-07-24 14:32:05.546 UTC [34] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
database     | 2020-07-24 14:32:05.588 UTC [35] LOG:  database system was shut down at 2020-07-24 14:32:05 UTC
database     | 2020-07-24 14:32:05.592 UTC [34] LOG:  database system is ready to accept connections
database     |  done
database     | server started
database     | CREATE DATABASE
database     |
database     |
database     | /usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/*
database     |
database     | 2020-07-24 14:32:05.748 UTC [34] LOG:  received fast shutdown request
database     | waiting for server to shut down....2020-07-24 14:32:05.750 UTC [34] LOG:  aborting any active transactions
database     | 2020-07-24 14:32:05.751 UTC [34] LOG:  background worker "logical replication launcher" (PID 41) exited with exit code 1
database     | 2020-07-24 14:32:05.751 UTC [36] LOG:  shutting down
database     | 2020-07-24 14:32:05.768 UTC [34] LOG:  database system is shut down
database     |  done
database     | server stopped
database     |
database     | PostgreSQL init process complete; ready for start up.
database     |
database     | 2020-07-24 14:32:05.862 UTC [1] LOG:  starting PostgreSQL 12.3 on x86_64-pc-linux-musl, compiled by gcc (Alpine 9.3.0) 9.3.0, 64-bit
database     | 2020-07-24 14:32:05.862 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
database     | 2020-07-24 14:32:05.862 UTC [1] LOG:  listening on IPv6 address "::", port 5432
database     | 2020-07-24 14:32:05.866 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
database     | 2020-07-24 14:32:05.910 UTC [45] LOG:  database system was shut down at 2020-07-24 14:32:05 UTC
database     | 2020-07-24 14:32:05.919 UTC [1] LOG:  database system is ready to accept connections
cli_1        | 2020-07-24 14:32:06.198383 I | [./neocargo-cli-shipment]
cli_1        | 2020-07-24 14:32:06.199659 I | Not enough arguments, expecting file and token
user-cli_1   | 2020-07-24 14:32:05.553226 I | Error creating user:  {"id":"go.micro.client","code":500,"detail":"service neocargo.service.user: not found","status":"Internal Server Error"}
user-cli_1   | {"id":"go.micro.client","code":500,"detail":"service neocargo.service.user: not found","status":"Internal Server Error"}
shipment_1   | 2020-07-24 14:32:05  file=v2@v2.9.1/service.go:200 level=info Starting [service] neocargo.service.shipment
shipment_1   | 2020-07-24 14:32:05  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:50051
shipment_1   | 2020-07-24 14:32:05  file=grpc/grpc.go:697 level=info Registry [mdns] Registering node: neocargo.service.shipment-1741495b-a5d2-41e6-9847-560329cc5101
vessel_1     | 2020-07-24 14:32:05  file=v2@v2.9.1/service.go:200 level=info Starting [service] neocargo.service.vessel
vessel_1     | 2020-07-24 14:32:05  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:50051
vessel_1     | 2020-07-24 14:32:05  file=grpc/grpc.go:697 level=info Registry [mdns] Registering node: neocargo.service.vessel-080f792b-9106-46de-ae27-e61c8c8e1a95
user_1       | --------------------------------------------------------
user_1       |  docker-compose-wait 2.7.3
user_1       | ---------------------------
user_1       | Starting with configuration:
user_1       |  - Hosts to be waiting for: [database:5432]
user_1       |  - Timeout before failure: 30 seconds
user_1       |  - TCP connection timeout before retry: 5 seconds
user_1       |  - Sleeping time before checking for hosts availability: 0 seconds
user_1       |  - Sleeping time once all hosts are available: 0 seconds
user_1       |  - Sleeping time between retries: 1 seconds
user_1       | --------------------------------------------------------
user_1       | Checking availability of database:5432
user_1       | Host database:5432 not yet available...
user_1       | Host database:5432 not yet available...
user_1       | Host database:5432 is now available!
user_1       | --------------------------------------------------------
user_1       | docker-compose-wait - Everything's fine, the application can now start!
user_1       | --------------------------------------------------------
user_1       | 2020-07-24 14:32:06.051395 I | conn: host=database user=postgres dbname=neocargo password=password sslmode=disable
user_1       | 2020-07-24 14:32:06.065657 I | sqlx connect done
user_1       | 2020-07-24 14:32:06.065737 I | DB connection OK: &{0xc000242900 postgres false 0xc00030b260}
user_1       | 2020-07-24 14:32:06  file=v2@v2.9.1/service.go:200 level=info Starting [service] neocargo.service.user
user_1       | 2020-07-24 14:32:06  file=grpc/grpc.go:864 level=info Server [grpc] Listening on [::]:50051
user_1       | 2020-07-24 14:32:06  file=grpc/grpc.go:697 level=info Registry [mdns] Registering node: neocargo.service.user-6d44822a-0fd1-4389-915e-e7ea8378b68b
neocargo_user-cli_1 exited with code 1
neocargo_cli_1 exited with code 1
```

You can stop all of your current containers by running:

```sh
# Teardown stack
$ make stop
```

### Testing

Test it all by running our CLI tool.

To run it through docker-compose, run:

- Shipment CLI tool

```sh
$ make run-cli

docker-compose run cli \
        ./neocargo-cli-shipment \
        shipment.json \
        eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ
Starting datastore ... done
Starting neocargo_vessel_1 ... done
Starting neocargo_shipment_1 ... done
2020-07-24 14:39:38.524923 I | [./neocargo-cli-shipment shipment.json eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9.TJVA95OrM7E2cBab30RMHrHDcEfxjoYZgeFONFh7HgQ]
2020-07-24 14:39:39.063664 I | Could not create a shipment: {"id":"go.micro.client","code":500,"detail":"error fetching vessel, returned nil","status":"Internal Server Error"}
make: *** [Makefile:35: run-cli] Error 1
```

- User CLI tool

```sh
$ make run-user-cli

docker-compose run user-cli \
        -- \
        --name="John Doe" \
        --email="john@foo.bar" \
        --password="test#test123" \
        --company="SpaceY"
Starting database ... done
Starting neocargo_user_1 ... done
Created user. Response:
2020-07-24 14:42:21.887597 I | Could not list users:  {"id":"go.micro.client","code":500,"detail":"must pass a pointer, not a value, to StructScan destination","status":"Internal Server Error"}
{"id":"go.micro.client","code":500,"detail":"must pass a pointer, not a value, to StructScan destination","status":"Internal Server Error"}
make: *** [Makefile:41: run-user-cli] Error 1
```