package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

type Storage struct {
	db *mongo.Client
}

func New(db *mongo.Client) *Storage {
	return &Storage{
		db: db,
	}
}

func (s *Storage) Create(collection *mongo.Collection, ctx context.Context, template map[string]interface{}) (map[string]interface{}, error) {
	req, err := collection.InsertOne(ctx, template)
	if err != nil {
		return nil, err
	}
	insertedId := req.InsertedID

	res := map[string]interface{}{
		"template": map[string]interface{}{
			"insertedId": insertedId,
		},
	}

	return res, nil
}

func (s *Storage) GetById(ctx context.Context, ID uuid.UUID) error {
	panic("Implement me")
}

func (s *Storage) GetByList(ctx context.Context) error {
	panic("Implement me")
}

func (s *Storage) Delete(ctx context.Context, ID uuid.UUID) error {
	panic("Implement me")
}
