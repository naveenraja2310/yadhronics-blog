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

func GetAllBlogs(ctx context.Context, limit, offset int64, search, category, status string) ([]models.Blogs, int64, error) {
	var blogs []models.Blogs

	// Define find options for pagination and sorting
	filter := bson.M{}

	if search != "" {
		filter["$or"] = []bson.M{
			{"title": bson.M{"$regex": search, "$options": "i"}},
			{"description": bson.M{"$regex": search, "$options": "i"}},
			{"category": bson.M{"$regex": search, "$options": "i"}},
		}
	}

	if category != "" {
		filter["category"] = category
	}

	if status != "" {
		filter["status"] = status
	}

	findOptions := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}}) // Sort by created_at field, descending

	// Apply limit if it's greater than 0
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}

	// Apply offset for pagination
	if offset > 0 {
		findOptions.SetSkip(int64(offset))
	}

	// Exclude "content" field from the result
	findOptions.SetProjection(bson.M{"content": 0})

	cursor, err := database.Blogs.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Decode the blogs from the cursor
	if err = cursor.All(ctx, &blogs); err != nil {
		return nil, 0, err
	}

	count, err := database.Blogs.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return blogs, count, nil
}

func GetBlogGroup(ctx context.Context, limit, offset int64) ([]bson.M, error) {
	matchStage := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "status", Value: "published"},
		}},
	}

	sortBeforeGroup := bson.D{
		{Key: "$sort", Value: bson.D{
			{Key: "created_at", Value: -1},
		}},
	}

	groupStage := bson.D{
		{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$category"},
			{Key: "blogs", Value: bson.D{
				{Key: "$push", Value: bson.D{
					{Key: "_id", Value: "$_id"},
					{Key: "title", Value: "$title"},
					{Key: "thumbnail", Value: "$thumbnail"},
					{Key: "description", Value: "$description"},
					{Key: "category", Value: "$category"},
					{Key: "type", Value: "$type"},
					{Key: "status", Value: "$status"},
					{Key: "created_at", Value: "$created_at"},
					{Key: "updated_at", Value: "$updated_at"},
				}},
			}},
			{Key: "count", Value: bson.D{{Key: "$sum", Value: 1}}},
		}},
	}

	projectStage := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "count", Value: 1},
			{Key: "blogs", Value: bson.D{
				{Key: "$slice", Value: bson.A{
					"$blogs",
					offset, limit,
				}},
			}},
		}},
	}

	sortStage := bson.D{
		{Key: "$sort", Value: bson.D{
			{Key: "_id", Value: 1}, // Sort by category name ascending
		}},
	}

	pipeline := mongo.Pipeline{matchStage, sortBeforeGroup, groupStage, projectStage, sortStage}

	cursor, err := database.Blogs.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []bson.M
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func GetAllCategories(ctx context.Context) ([]string, error) {
	var categories []string

	distinctValues, err := database.Blogs.Distinct(ctx, "category", bson.M{})
	if err != nil {
		return nil, err
	}
	for _, value := range distinctValues {
		if strValue, ok := value.(string); ok {
			categories = append(categories, strValue)
		}
	}

	return categories, nil
}
