package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// https://medium.com/better-programming/building-a-restful-api-with-go-and-mongodb-93e59cbbee88
// https://www.mongodb.com/blog/post/quick-start-golang--mongodb--starting-and-setup

// GetDbClient ... returns a new db connection
func getDbClient() (*mongo.Client, error) {

	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, fmt.Errorf("connecting to db: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = testDbConnection(ctx, c)
	if err != nil {
		return nil, fmt.Errorf("pinging db: %w", err)
	}

	return c, nil

}

func testDbConnection(ctx context.Context, c *mongo.Client) error {
	err := c.Ping(ctx, readpref.Primary())
	if err != nil {
		return fmt.Errorf("testing db connection: %w", err)
	}
	return nil
}

func getDbCollection(CollectionName string, DbName string) (*mongo.Collection, error) {
	c, err := getDbClient()
	if err != nil {
		return nil, fmt.Errorf("get db collection: %w", err)
	}

	coll := c.Database(DbName).Collection(CollectionName)

	return coll, nil
}
