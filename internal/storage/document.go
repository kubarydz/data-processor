package storage

import (
	"time"

	"github.com/kubarydz/data-processor/internal/domain/user"
)

type userDocument struct {
	Id        string    `bson:"_id"`
	Name      string    `bson:"name"`
	Timestamp time.Time `bson:"timestamp"`
}

func newUserDocument(user *user.User) *userDocument {
	return &userDocument{
		Id:   user.Id,
		Name: user.Name,
	}
}
