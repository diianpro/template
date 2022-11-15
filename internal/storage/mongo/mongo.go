package mongo

import (
	"context"
	"fmt"
	"github.com/diianpro/template/internal/domain"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"unsafe"
)

type Storage struct {
	db *mongo.Database
}

func New(cli *mongo.Client) *Storage {
	return &Storage{
		db: cli.Database("templates"),
	}
}

func (s *Storage) Create(ctx context.Context, template []byte) (string, error) {
	ID := uuid.New().String()
	_, err := s.db.Collection("template").InsertOne(ctx, bson.D{
		{"id", ID},
		{"file", template},
	},
	)
	if err != nil {
		return "", fmt.Errorf("create: error: %w", err)
	}

	return ID, nil
}

func (s *Storage) Delete(ctx context.Context, id uuid.UUID) error {
	file, err := s.db.Collection("template").DeleteOne(ctx, bson.M{
		"id": id.String(),
	})
	if err != nil {
		return fmt.Errorf("delete by id: error: %w", err)

	}
	if file.DeletedCount == 0 {
		return fmt.Errorf("no file were deleted: %w", err)
	}
	return nil
}

func (s *Storage) GetByID(ctx context.Context, id uuid.UUID) ([]byte, error) {
	var r = struct {
		File []byte `bson:"file"`
	}{}
	res := s.db.Collection("template").FindOne(ctx, bson.M{
		"id": id.String(),
	}, options.FindOne().SetProjection(bson.D{{"file", 1}, {"_id", 0}}))
	if res.Err() != nil {
		return nil, fmt.Errorf("get by: find error: %w", res.Err())
	}
	err := res.Decode(&r)
	if err != nil {
		return nil, fmt.Errorf("get by id: decode error: %w", err)
	}
	return r.File, nil
}

func (s *Storage) GetAll(ctx context.Context, limit int64, offset int64) (*domain.Templates, error) {
	type resp = struct {
		ID   string
		File []byte `bson:"file"`
	}

	var result []resp

	findOptions := options.Find().SetProjection(bson.D{{"file", 1}, {"id", 1}, {"_id", 0}})
	findOptions.SetLimit(limit)
	findOptions.SetSkip(offset)

	crs, err := s.db.Collection("template").Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, fmt.Errorf("not find: error : %w", err)
	}

	for crs.Next(ctx) {
		var elem resp
		err = crs.Decode(&elem)
		if err != nil {
			return nil, fmt.Errorf("decode elem: error : %w", err)
		}
		result = append(result, elem)
	}
	if err = crs.Err(); err != nil {
		return nil, fmt.Errorf("not find all: error : %w", err)
	}
	ret := (*domain.Templates)(unsafe.Pointer(&result))

	return ret, err
}
