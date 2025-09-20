package service

import (
	"context"
	"time"
	"yadhronics-blog/database"
	"yadhronics-blog/models"

	"go.mongodb.org/mongo-driver/mongo"
)

func CreateBlog(ctx context.Context, blog models.Blogs) (*mongo.InsertOneResult, error) {
	blog.CreatedAt = time.Now()

	result, err := database.Blogs.InsertOne(ctx, blog)
	if err != nil {
		return nil, err
	}

	return result, nil
}
