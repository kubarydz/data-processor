package user

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type Repository interface {
	Save(ctx context.Context, user *User) error
}

type Service struct {
	repo   Repository
	reader *kafka.Reader
}

func NewService(reader *kafka.Reader, repo Repository) Service {
	return Service{
		repo:   repo,
		reader: reader,
	}
}

func (s *Service) Run() {
	for {
		ctx := context.Background()
		msg, err := s.reader.ReadMessage(context.Background())
		if err != nil {
			log.Println("Failed to read msg: ", err)
			continue
		}

		log.Println("Received: ", string(msg.Key), string(msg.Value))

		user := NewUser(string(msg.Key), string(msg.Value))
		err = s.repo.Save(ctx, user)
		if err != nil {
			log.Println("Failed to save user: ", err)
		}

	}

}
