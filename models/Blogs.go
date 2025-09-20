package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blogs struct {
	Id          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Thumbnail   string             `bson:"thumbnail" json:"thumbnail"`
	Description string             `bson:"description" json:"description"`
	Category    string             `bson:"category" json:"category"`
	Type        string             `bson:"type" json:"type"`
	Content     []DocNode          `bson:"content" json:"content"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}
type DocNode struct {
	Type    string     `bson:"type,omitempty" json:"type,omitempty"`
	Attrs   *NodeAttrs `bson:"attrs,omitempty" json:"attrs,omitempty"`
	Content []DocNode  `bson:"content,omitempty" json:"content,omitempty"`
	Text    string     `bson:"text,omitempty" json:"text,omitempty"`
	Marks   []Mark     `bson:"marks,omitempty" json:"marks,omitempty"`
}

type NodeAttrs struct {
	TextAlign *string `bson:"textAlign,omitempty" json:"textAlign,omitempty"`
	Level     *int    `bson:"level,omitempty" json:"level,omitempty"`
	Href      *string `bson:"href,omitempty" json:"href,omitempty"`
	Target    *string `bson:"target,omitempty" json:"target,omitempty"`
	Rel       *string `bson:"rel,omitempty" json:"rel,omitempty"`
	Class     *string `bson:"class,omitempty" json:"class,omitempty"`
	Color     *string `bson:"color,omitempty" json:"color,omitempty"`
	Src       *string `bson:"src,omitempty" json:"src,omitempty"`
	Alt       *string `bson:"alt,omitempty" json:"alt,omitempty"`
	Title     *string `bson:"title,omitempty" json:"title,omitempty"`
	Width     *int    `bson:"width,omitempty" json:"width,omitempty"`
	Height    *int    `bson:"height,omitempty" json:"height,omitempty"`
	Checked   *bool   `bson:"checked,omitempty" json:"checked,omitempty"`
	Language  *string `bson:"language,omitempty" json:"language,omitempty"`
}

type Mark struct {
	Type  string         `bson:"type,omitempty" json:"type,omitempty"`
	Attrs map[string]any `bson:"attrs,omitempty" json:"attrs,omitempty"`
}

func (b Blogs) TableName() string {
	return "blogs"
}
