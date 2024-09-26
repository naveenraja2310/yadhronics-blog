package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskHistory struct {
	Id         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Task       Task               `bson:"Task,omitempty" json:"Task,omitempty"`
	FromStatus TaskStatus         `bson:"FromStatus,omitempty" json:"FromStatus,omitempty"`
	ToStatus   TaskStatus         `bson:"ToStatus,omitempty" json:"ToStatus,omitempty"`
	ChangedBy  User               `bson:"ChangedBy,omitempty" json:"ChangedBy,omitempty"`
	Remark     string             `bson:"Remark,omitempty" json:"Remark,omitempty"`
	CreatedAt  time.Time          `bson:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	UpdatedAt  time.Time          `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"`
}
