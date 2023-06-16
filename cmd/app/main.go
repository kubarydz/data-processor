package main

import (
	"log"

	"github.com/kubarydz/data-processor/internal/domain/user"
	"github.com/kubarydz/data-processor/internal/storage"
	"github.com/segmentio/kafka-go"
)

func main() {

	kafkaReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "random-data",
		GroupID:  "cons-group",
		MinBytes: 10e1,
		MaxBytes: 10e6,
	})

	defer kafkaReader.Close()

	collection, err := storage.ProvideCollection()
	if err != nil {
		log.Println("Failed to connect to db: ", err)
		panic(err)
	}
	userRepo := storage.ProvideRepository(collection)

	userService := user.NewService(kafkaReader, userRepo)
	userService.Run()
}
