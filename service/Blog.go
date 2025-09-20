package service

import (
	"context"
	"errors"
	"time"
	"yadhronics-blog/database"
	"yadhronics-blog/models"
	"yadhronics-blog/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateBlog(ctx context.Context, blog models.Blogs) (*mongo.InsertOneResult, error) {
	blog.Status = "draft"
	blog.CreatedAt = time.Now()

	result, err := database.Blogs.InsertOne(ctx, blog)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func UpdateBlog(ctx context.Context, blog models.Blogs, id primitive.ObjectID) (*models.Blogs, error) {
	if !utils.CheckIfExistsByID(ctx, database.Blogs, id) {
		return nil, errors.New("the given id is invalid")
	}

	updateFields := bson.M{"updated_at": time.Now()}

	updateFields["title"] = blog.Title
	updateFields["thumbnail"] = blog.Thumbnail
	updateFields["description"] = blog.Description
	updateFields["category"] = blog.Category
	updateFields["type"] = blog.Type
	updateFields["content"] = blog.Content
	updateFields["status"] = blog.Status

	update := bson.M{"$set": updateFields}
	result, err := database.Blogs.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("no document found with the given id")
	}

	updatedBlog, err := GetBlogByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return updatedBlog, nil
}

func GetBlogByID(ctx context.Context, id primitive.ObjectID) (*models.Blogs, error) {
	var blog models.Blogs

	if !utils.CheckIfExistsByID(ctx, database.Blogs, id) {
		return nil, errors.New("the given id is invalid")
	}

	err := database.Blogs.FindOne(ctx, bson.M{"_id": id}).Decode(&blog)
	if err != nil {
		return nil, err
	}

	return &blog, nil
}

func DeleteBlog(ctx context.Context, id primitive.ObjectID) error {
	if !utils.CheckIfExistsByID(ctx, database.Blogs, id) {
		return errors.New("the given id is invalid")
	}

	result, err := database.Blogs.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("no document found with the given id")
	}

	return nil
}

func GetAllBlogs(ctx context.Context, limit, offset int64) ([]models.Blogs, int64, error) {
	var blogs []models.Blogs

	// Define find options for pagination and sorting
	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}) // Sort by created_at field, descending

	// Apply limit if it's greater than 0
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}

	// Apply offset for pagination
	if offset > 0 {
		findOptions.SetSkip(int64(offset))
	}

	cursor, err := database.Blogs.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Decode the users from the cursor
	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, 0, err
	}

	count, err := database.Blogs.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}
