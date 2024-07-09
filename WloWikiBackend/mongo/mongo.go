package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client wraps the MongoDB client
type Client struct {
	*mongo.Client
}

// NewClient initializes and returns a new MongoDB client
func NewClient(uri string) (*Client, error) {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	// Check the connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}

	return &Client{client}, nil
}
