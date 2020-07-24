# Use multi-stage build to avoid storing GitHub token inside Docker image
FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ARG GITHUB_TOKEN

ENV GO111MODULE=on
# Required to access go mod from private neocargo repo
ENV GOPRIVATE github.com/haxorbit/neocargo

RUN echo "GH_TOK: $GITHUB_TOKEN"

RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/haxorbit/".insteadOf "https://github.com/haxorbit/"

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o neocargo-cli-user *.go

# Run container
FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/neocargo-cli-user .

ENTRYPOINT ["./neocargo-cli-user"]