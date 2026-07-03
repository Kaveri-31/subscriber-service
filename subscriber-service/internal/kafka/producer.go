package kafka

import (
	"context"
	"log"

	kafkago "github.com/segmentio/kafka-go"

	"subscriber-service/internal/config"
)

var Writer *kafkago.Writer

// InitProducer initializes the Kafka producer
func InitProducer() error {

	cfg := config.Get()

	Writer = &kafkago.Writer{
		Addr: kafkago.TCP(cfg.Kafka.Broker),

		Topic: cfg.Kafka.Topic,

		Balancer: &kafkago.LeastBytes{},
	}

	log.Println("Kafka Producer Initialized")

	return nil
}

// Publish sends a message to Kafka
func Publish(message string) {

	if Writer == nil {
		log.Println("Kafka Producer is not initialized")
		return
	}

	err := Writer.WriteMessages(
		context.Background(),
		kafkago.Message{
			Value: []byte(message),
		},
	)

	if err != nil {
		log.Println("Kafka Publish Error:", err)
		return
	}

	log.Println("Published:", message)
}

// CloseProducer closes the Kafka producer
func CloseProducer() {

	if Writer != nil {

		if err := Writer.Close(); err != nil {
			log.Println("Error closing Kafka Producer:", err)
		}
	}
}

// Liveness Check
func IsConnected() bool {

	return Writer != nil
}
