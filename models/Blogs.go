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
	Type        string             `json:"type"`
	Content     []DocNode          `json:"content"`
	Status      string             `bson:"status" json:"status"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

type DocNode struct {
	Type    string     `json:"type"`
	Attrs   *NodeAttrs `json:"attrs,omitempty"`
	Content []DocNode  `json:"content,omitempty"`
	Text    string     `json:"text,omitempty"`
	Marks   []Mark     `json:"marks,omitempty"`
}

type NodeAttrs struct {
	TextAlign *string `json:"textAlign,omitempty"`
	Level     *int    `json:"level,omitempty"`
	Href      *string `json:"href,omitempty"`
	Target    *string `json:"target,omitempty"`
	Rel       *string `json:"rel,omitempty"`
	Class     *string `json:"class,omitempty"`
	Color     *string `json:"color,omitempty"`
	Src       *string `json:"src,omitempty"`
	Alt       *string `json:"alt,omitempty"`
	Title     *string `json:"title,omitempty"`
	Width     *int    `json:"width,omitempty"`
	Height    *int    `json:"height,omitempty"`
	Checked   *bool   `json:"checked,omitempty"`
	Language  *string `json:"language,omitempty"`
}

type Mark struct {
	Type  string         `json:"type"`
	Attrs map[string]any `json:"attrs,omitempty"`
}

func (b Blogs) TableName() string {
	return "blogs"
}
