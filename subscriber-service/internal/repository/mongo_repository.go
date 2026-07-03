package repository

import (
	"context"
	"errors"
	"time"

	"subscriber-service/internal/database"
	"subscriber-service/internal/model"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Returns the MongoDB collection
func getCollection() *mongo.Collection {
	return database.GetCollection()
}

// Create Subscriber
func CreateSubscriber(subscriber model.Subscriber) (*model.Subscriber, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	collection := getCollection()

	if collection == nil {
		return nil, errors.New("mongodb collection is nil")
	}

	result, err := collection.InsertOne(
		ctx,
		subscriber,
	)

	if err != nil {
		return nil, err
	}

	// Store generated MongoDB ObjectID
	objectID, ok := result.InsertedID.(bson.ObjectID)

	if ok {
		subscriber.ID = objectID
	}

	return &subscriber, nil
}

// Get All Subscribers
func GetAllSubscribers() ([]model.Subscriber, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	collection := getCollection()

	if collection == nil {
		return nil, errors.New("mongodb collection is nil")
	}

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	subscribers := make([]model.Subscriber, 0)

	for cursor.Next(ctx) {

		var subscriber model.Subscriber

		if err := cursor.Decode(&subscriber); err != nil {
			return nil, err
		}

		subscribers = append(
			subscribers,
			subscriber,
		)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return subscribers, nil
}

// Get Subscriber By ID
func GetSubscriberByID(id string) (*model.Subscriber, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	collection := getCollection()

	if collection == nil {
		return nil, errors.New("mongodb collection is nil")
	}

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid subscriber id")
	}

	var subscriber model.Subscriber

	err = collection.FindOne(
		ctx,
		bson.M{
			"_id": objectID,
		},
	).Decode(&subscriber)

	if err != nil {
		return nil, err
	}

	return &subscriber, nil
}

// Update Subscriber
func UpdateSubscriber(
	id string,
	subscriber model.Subscriber,
) (*model.Subscriber, error) {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	collection := getCollection()

	if collection == nil {
		return nil, errors.New("mongodb collection is nil")
	}

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return nil, errors.New("invalid subscriber id")
	}

	update := bson.M{
		"$set": bson.M{
			"imsi":   subscriber.IMSI,
			"msisdn": subscriber.MSISDN,
			"plan":   subscriber.Plan,
			"status": subscriber.Status,
		},
	}

	_, err = collection.UpdateOne(
		ctx,
		bson.M{
			"_id": objectID,
		},
		update,
	)

	if err != nil {
		return nil, err
	}

	return GetSubscriberByID(id)
}

// Delete Subscriber
func DeleteSubscriber(id string) error {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)
	defer cancel()

	collection := getCollection()

	if collection == nil {
		return errors.New("mongodb collection is nil")
	}

	objectID, err := bson.ObjectIDFromHex(id)

	if err != nil {
		return errors.New("invalid subscriber id")
	}

	_, err = collection.DeleteOne(
		ctx,
		bson.M{
			"_id": objectID,
		},
	)

	if err != nil {
		return err
	}

	return nil
}
