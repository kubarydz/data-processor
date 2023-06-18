package main

import (
	"context"
	"log"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/segmentio/kafka-go"
)

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "random-data",
		Balancer: &kafka.LeastBytes{},
	})
	writer.AllowAutoTopicCreation = true

	defer writer.Close()

	for {
		time.Sleep(1 * time.Second)
		data := struct {
			ID    string
			Value string
		}{
			ID:    faker.UUIDHyphenated(),
			Value: faker.LastName(),
		}

		msg := kafka.Message{
			Key:   []byte(data.ID),
			Value: []byte(data.Value),
		}

		err := writer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Println("Failed to write msg: ", err)
			continue
		}

		log.Println("Published: ", data)

	}

}
