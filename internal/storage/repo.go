package storage

import (
	"context"

	"github.com/kubarydz/data-processor/internal/domain/user"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository struct {
	collection *mongo.Collection
}

func ProvideRepository(collection *mongo.Collection) Repository {
	return Repository{collection}
}

func (r Repository) Save(ctx context.Context, user *user.User) error {
	doc := newUserDocument(user)
	_, err := r.collection.InsertOne(ctx, doc)
	return err
}
