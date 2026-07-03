package service

import (
	"encoding/json"
	"log"

	"subscriber-service/internal/kafka"
	"subscriber-service/internal/model"
	"subscriber-service/internal/repository"
)

// Get all subscribers
func GetAllSubscribers() ([]model.Subscriber, error) {

	return repository.GetAllSubscribers()
}

// Get subscriber by ID
func GetSubscriberByID(id string) (*model.Subscriber, error) {

	return repository.GetSubscriberByID(id)
}

// Create Subscriber
func CreateSubscriber(subscriber model.Subscriber) (*model.Subscriber, error) {

	createdSubscriber, err := repository.CreateSubscriber(subscriber)

	if err != nil {
		return nil, err
	}

	event := map[string]interface{}{
		"event":      "subscriber.created",
		"subscriber": createdSubscriber,
	}

	eventJSON, err := json.Marshal(event)

	if err == nil {
		kafka.Publish(string(eventJSON))
		log.Println("Kafka Event Published : subscriber.created")
	}

	return createdSubscriber, nil
}

// Update Subscriber
func UpdateSubscriber(id string, subscriber model.Subscriber) (*model.Subscriber, error) {

	updatedSubscriber, err := repository.UpdateSubscriber(id, subscriber)

	if err != nil {
		return nil, err
	}

	event := map[string]interface{}{
		"event":      "subscriber.updated",
		"subscriber": updatedSubscriber,
	}

	eventJSON, err := json.Marshal(event)

	if err == nil {
		kafka.Publish(string(eventJSON))
		log.Println("Kafka Event Published : subscriber.updated")
	}

	return updatedSubscriber, nil
}

// Delete Subscriber
func DeleteSubscriber(id string) error {

	err := repository.DeleteSubscriber(id)

	if err != nil {
		return err
	}

	event := map[string]interface{}{
		"event": "subscriber.deleted",
		"id":    id,
	}

	eventJSON, err := json.Marshal(event)

	if err == nil {
		kafka.Publish(string(eventJSON))
		log.Println("Kafka Event Published : subscriber.deleted")
	}

	return nil
}
