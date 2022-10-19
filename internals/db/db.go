package db

import (
	"context"
	"os"
	"time"

	"github.com/Napigo/npglogger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DS struct {
	*mongo.Client
}

var DataSource *DS

func ConnectDB() {
	URI := os.Getenv("MONGO_DATABASE_URI")

	client, err := mongo.NewClient(options.Client().ApplyURI(URI))
	if err != nil {
		npglogger.Fatalf("Failed to create Mongo DB client with URI of %s", URI)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		npglogger.Fatalf("Failed to create a connection stream to MongoDB with URI of %s", URI)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		npglogger.Fatalf("Failed to establish connection with mongodb, URI of %s", URI)
	} else {
		npglogger.Infof("MongoDB Connected")
	}
	DataSource = &DS{client}
}

func DisconnectDB(c *mongo.Client, ctx context.Context) {
	if err := c.Disconnect(ctx); err != nil {
		npglogger.Fatalf("Failed to disconnect from mongodb with URI of %s", os.Getenv("MONGO_DATABASE_URI"))
	}
}
