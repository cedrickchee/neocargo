# This is a multi-stage Dockerfile to make our containers smaller so
# they can deploy faster, scale faster, and take up less storage space.

# Build container
# Fetch dependencies, use the Go runtime image as the base image
FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download
# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o neocargo-service-shipment *.go

# Run container
# Take the binary from build container and then run binary,
# using the Alpine base image
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app
COPY --from=builder /app/neocargo-service-shipment .

CMD ["./neocargo-service-shipment"]
