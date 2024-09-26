package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"Name,omitempty" json:"Name,omitempty"`
	Description string             `bson:"Description,omitempty" json:"Description,omitempty"`
	FromUser    User               `bson:"FromUser,omitempty" json:"FromUser,omitempty"`
	ToUser      User               `bson:"ToUser,omitempty" json:"ToUser,omitempty"`
	TaskStatus  TaskStatus         `bson:"TaskStatus,omitempty" json:"TaskStatus,omitempty"`
	TaskType    TaskType           `bson:"TaskType,omitempty" json:"TaskType,omitempty"`
	Remark      string             `bson:"Remark,omitempty" json:"Remark,omitempty"`
	CreatedAt   time.Time          `bson:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time          `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"`
}
