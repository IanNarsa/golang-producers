package main

import (
	"context"
	"fmt"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	// Configure Kafka producer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "my-kafka-topic",
		Balancer: &kafka.LeastBytes{}, // Optional load balancing strategy
	})

	defer w.Close()

	// Produce messages
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Message %d", i)
		err := w.WriteMessages(context.Background(), kafka.Message{
			Value: []byte(message),
		})
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("Produced message: %s\n", message)
	}
}
