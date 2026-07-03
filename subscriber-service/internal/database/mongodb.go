package database

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"

	"subscriber-service/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var (
	Client     *mongo.Client
	Collection *mongo.Collection
)

// ConnectMongoDB establishes a connection to MongoDB
func ConnectMongoDB() {

	cfg, err := config.Load("configs/config.yaml")

	if err != nil {
		logrus.WithError(err).Fatal("Unable to load configuration")
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	client, err := mongo.Connect(
		options.Client().
			ApplyURI(cfg.MongoDB.URI),
	)

	if err != nil {
		logrus.WithError(err).Fatal("MongoDB Connection Failed")
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		logrus.WithError(err).Fatal("MongoDB Ping Failed")
	}

	Client = client

	Collection = Client.
		Database(cfg.MongoDB.Database).
		Collection(cfg.MongoDB.Collection)

	logrus.Info("MongoDB Connected Successfully")
}

// GetCollection returns the MongoDB collection
func GetCollection() *mongo.Collection {

	return Collection
}

// DisconnectMongoDB closes the MongoDB connection
func DisconnectMongoDB() {

	if Client == nil {
		return
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	if err := Client.Disconnect(ctx); err != nil {
		logrus.WithError(err).Error("Error disconnecting MongoDB")
		return
	}

	logrus.Info("MongoDB Connection Closed")
}

// Liveness Check
func IsConnected() bool {

	return Collection != nil
}
