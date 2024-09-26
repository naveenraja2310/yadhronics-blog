package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskType struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Type        string             `bson:"Type,omitempty" json:"Type,omitempty"`
	Description string             `bson:"Description,omitempty" json:"Description,omitempty"`
	CreatedAt   time.Time          `bson:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time          `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"`
}

func (t TaskType) TableName() string {
	return "tasktype"
}
