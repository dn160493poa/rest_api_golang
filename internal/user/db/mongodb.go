package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"restApi/internal/user"
	"restApi/pkg/logging"
)

type db struct {
	collection *mongo.Collection
	logger     *logging.Logger
}

func (d *db) Create(ctx context.Context, user user.User) (string, error) {
	d.collection.InsertOne(ctx, user)
	panic("implement me")
}

func (d *db) Find(ctx context.Context, id string) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (d *db) Update(ctx context.Context, user user.User) error {
	//TODO implement me
	panic("implement me")
}

func (d *db) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewStorage(database *mongo.Database, collection string, logger *logging.Logger) user.Storage {
	return &db{
		collection: database.Collection(collection),
		logger:     logger,
	}
}
