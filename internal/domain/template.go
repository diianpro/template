package domain

import "github.com/google/uuid"

type Template struct {
	ID uuid.UUID `bson:"_id" json:"id"`
}
