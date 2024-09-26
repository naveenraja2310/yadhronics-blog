package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name        string             `bson:"Name,omitempty" json:"Name,omitempty"`
	Email       string             `bson:"Email,omitempty" json:"Email,omitempty"`
	PhoneNumber string             `bson:"PhoneNumber,omitempty" json:"PhoneNumber,omitempty"`
	IsVerified  bool               `bson:"IsVerified" json:"IsVerified"`
	IsAdmin     bool               `bson:"IsAdmin" json:"IsAdmin"`
	Photo       string             `bson:"Photo,omitempty" json:"Photo,omitempty"`
	CreatedAt   time.Time          `bson:"CreatedAt,omitempty" json:"CreatedAt,omitempty"`
	UpdatedAt   time.Time          `bson:"UpdatedAt,omitempty" json:"UpdatedAt,omitempty"`
}

func (t User) TableName() string {
	return "user"
}
