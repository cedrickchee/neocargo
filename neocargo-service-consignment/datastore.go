// Create the master session/connection

package main

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient creates a connection, using a given connection string
func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	conn, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))

	// 'Pinging' the connection to check it's correct and that the datastore is
	// available.
	if err := conn.Ping(ctx, nil); err != nil {
		// Basic retry logic
		if retry >= 3 {
			// If it exceeds three retries, we let the error bubble upwards to
			// be handled.
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)

		// Calling itself again if it can't connect
		return CreateClient(ctx, uri, retry)
	}

	return conn, err
}
