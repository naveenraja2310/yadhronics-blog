package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskStatus struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"Name,omitempty" json:"Name,omitempty"`
	Description string             `bson:"Description,omitempty" json:"Description,omitempty"`
	CreatedAt   time.Time          `bson:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time          `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"`
}

func (t TaskStatus) TableName() string {
	return "taskstatus"
}
