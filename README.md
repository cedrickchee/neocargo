# neoCargo Microservices in Go

neoCargo is a shipping container management platform.

This project is an example of multiple microservices implementation in Go in a [**monorepo**](https://en.wikipedia.org/wiki/Monorepo).

(_Demo ([asciinema](https://asciinema.org/) coming soon._)

The neoCargo backend consists of 3 microservices:

- [Shipments](./neocargo-service-shipment)
- [Vessels](./neocargo-service-vessel)
- [Users + Authentication](./neocargo-service-user)

## Tech Stack

- [Protocol Buffers](https://developers.google.com/protocol-buffers) and [gRPC](https://grpc.io/) as transport protocol
- [go-micro](https://micro.mu/)
- Docker container ([Alpine Linux](https://alpinelinux.org/about/) as base image)
- [Docker Compose](https://docs.docker.com/compose/)
- PostgreSQL or MongoDB database
- User authentication with [JWT](https://jwt.io/)
- [Google Cloud](https://cloud.google.com/)
- [Kubernetes](https://kubernetes.io/)
- [NATS](https://nats.io/)
- [CircleCI](https://circleci.com/)
- [Terraform](https://www.terraform.io/)

## System Architecture

(_Note: incomplete._)

A good architecture is when services are decoupled. We will achieve that with
event driven architecture or pubsub.

We need to integrate the NATS broker plug-in into our services.

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

### Local Development

**Build and run Docker Compose stack**

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

- Create a shipment

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

- Create a user

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

### Deployment

#### Terraform a Cluster on Google Cloud

We will create a cloud environment to host our services. We will be using
Terraform to provision and manage our cloud cluster on Google Cloud.

Steps:

- Create your [Google Cloud](http://console.cloud.google.com/) project.
- Go to "IAM & Admin" tab in Google Cloud console and [create a new service account key](https://cloud.google.com/iam/docs/creating-managing-service-account-keys). Make sure you select "JSON" for "Key type".
- Modify the configurations in our [Infra-as-Code project](./infra).
- Move the service key you created earlier into the project root and name it `gcloud-service-key.json`.

Next, create a new cluster:

```sh
$ terraform init

# View deployment plan
$ terraform plan

# Apply the changes
$ terraform apply
```

Once it's done, see your new cluster. Go to Google Cloud console and look for Kubernetes Engine (GKE).
Next, deploy our containers to the cluster.

#### Google Kubernetes Engine (GKE)

Set-up and deploy containers into cluster using GKE.

Steps:

- Ensure you have the [kubectl cli installed locally](https://cloud.google.com/kubernetes-engine/docs/quickstart#local-shell):

```
$ gcloud components install kubectl
```

Usually, you'd deploy a PostgreSQL/MongoDB instance, or database instance along
with every service, for complete separation.

Then we deploy our services, shipment-service, vessel-service, and user-service.

**MongoDB containers**

I created three Kubernetes deployment files for MongoDB:
- [storage](./deployment/mongodb-deployment.yml)
- [stateful set](./deployment/mongodb-deployment.yml)
- [our service](./deployment/service.yml)

```sh
$ kubectl create -f ./deployment/mongodb-ssd.yml
$ kubectl create -f ./deployment/mongodb-deployment.yml
$ kubectl create -f ./deployment/mongodb-service.yml
```

The result of this will be a replicated set of MongoDB containers, with stateful storage and a service exposing the datastore across our other pods.

**Vessel service**

I have created a [deployment file for vessel service](./deployment/vessel-service-deployment.yml).

We're pushing and pulling our Docker image from a private [Container Registry](https://cloud.google.com/container-registry/docs/pushing-and-pulling).

```sh
$ docker build -t asia.gcr.io/neocargo/vessel-service:latest .
$ gcloud docker -- push asia.gcr.io/neocargo/vessel-service:latest
```

Then, deploy vessel-service to our cluster.

```sh
$ kubectl create -f ./deployments/vessel-service-deployment.yml
$ kubectl create -f ./deployments/vessel-service.yml
```

Do the same for our other services.

**Deploy Micro**

[Deployment](./deployment/micro-deployment.yml) file.

In our service here, we expose an external load balancer, with an IP address out to the public.

Run `$ kubectl get services` to get a public IP address.

After all that's deployed, make a service call to micro container:

```sh
$ curl localhost/rpc -XPOST -d '{
    "request": {
        "name": "test",
        "capacity": 100,
        "max_weight": 200000,
        "available": true
    },
    "method": "VesselService.Create",
    "service": "vessel"
}' -H 'Content-Type: application/json'
```

Here, our gRPC services, being proxied and converted to a web friendly format, using a sharded, MongoDB instance.

### Continuous Integration (CI)

We will set up all of this into a CI process to manage our deployments.

**Set up CircleCI with our services**

First, sign-up and create a new project in CircleCI.

Next, change your build configuration (let's use our shipment-service for this).

Open [shipment-service build configuration](./neocargo-service-shipment/.circleci/config.yml)
in your editor. In order to make this work, we need Google Cloud service account
key, such as the one we created in ["Deployment" section](#Deployment) of this
README, and we need to encode this into base64 and store it as an environment
variable within our CircleCI build project settings. You can read more about
environment variables and how to use them on [CircleCI 2.0 here](https://circleci.com/docs/2.0/env-vars/).

That's it. We have CI for one of our services. For a production-grade service,
I recommend you run your tests first, before your deploy step.
