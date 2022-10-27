package mongo

import (
	"context"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"html/template"
	"log"
)

type Storage struct {
	db *mongo.Client
}

func New(db *mongo.Client) *Storage {
	db.Database("templates").Collection("template")
	return &Storage{
		db: db,
	}
}

func (s *Storage) Create(ctx context.Context, template []byte) (uuid.UUID, error) {
	req, err := s.db.Database("templates").Collection("template").InsertOne(ctx, bson.D{
		{"file", template},
	},
	)
	if err != nil {
		log.Println("There was an err in trying to migrate the data into the database")
	}

	insertedId := req.InsertedID.(uuid.UUID)

	return insertedId, nil
}

func (s *Storage) GetById(ctx context.Context, ID uuid.UUID) (template.Template, error) {
	var res template.Template
	err := s.db.Database("templates").Collection("template").FindOne(ctx, bson.M{
		"_id": ID,
	}).Decode(&res)
	if err != nil {
		log.Println("There was an error in trying to find by ID the data in the database")

	}
	return res, nil
}

func (s *Storage) GetByList(ctx context.Context) error {
	panic("Implement me")
}

func (s *Storage) Delete(ctx context.Context, ID uuid.UUID) error {
	_, err := s.db.Database("templates").Collection("template").DeleteOne(ctx, bson.M{
		"_id": ID,
	})
	if err != nil {
		log.Println("There was an error in trying to find by ID the data in the database")

	}
	return nil
}
