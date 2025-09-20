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

func CreateTaskType(ctx context.Context, tasktype models.TaskType) (*mongo.InsertOneResult, error) {
	tasktype.CreatedAt = time.Now()

	result, err := database.TaskType.InsertOne(ctx, tasktype)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTaskTypeByID(ctx context.Context, id primitive.ObjectID) (*models.TaskType, error) {
	var tasktype models.TaskType

	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	err := database.TaskType.FindOne(ctx, bson.M{"_id": id}).Decode(&tasktype)
	if err != nil {
		return nil, err
	}

	return &tasktype, nil
}

func GetAllTaskType(ctx context.Context, limit, offset int) ([]models.TaskType, int64, error) {
	var tasktype []models.TaskType

	// Define find options for pagination and sorting
	findOptions := options.Find().SetSort(bson.D{{Key: "CreatedAt", Value: -1}}) // Sort by created_at field, descending

	// Apply limit if it's greater than 0
	if limit > 0 {
		findOptions.SetLimit(int64(limit))
	}

	// Apply offset for pagination
	if offset > 0 {
		findOptions.SetSkip(int64(offset))
	}

	cursor, err := database.TaskType.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Decode the users from the cursor
	if err = cursor.All(ctx, &tasktype); err != nil {
		return nil, 0, err
	}

	count, err := database.TaskType.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return tasktype, count, nil
}

func UpdateTaskType(ctx context.Context, id primitive.ObjectID, tasktype models.TaskType) (*mongo.UpdateResult, error) {

	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	updateFields := bson.M{"UpdatedAt": time.Now()}

	updateFields["Type"] = tasktype.Type
	updateFields["Description"] = tasktype.Description

	update := bson.M{"$set": updateFields}
	result, err := database.TaskType.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteTaskType(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {

	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	result, err := database.TaskType.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}
