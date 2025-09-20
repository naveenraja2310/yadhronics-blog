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

func CreateTaskStatus(ctx context.Context, taskstatus models.TaskStatus) (*mongo.InsertOneResult, error) {
	taskstatus.CreatedAt = time.Now()

	result, err := database.TaskStatus.InsertOne(ctx, taskstatus)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTaskStatusByID(ctx context.Context, id primitive.ObjectID) (*models.TaskStatus, error) {
	var taskstatus models.TaskStatus

	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}
	err := database.TaskStatus.FindOne(ctx, bson.M{"_id": id}).Decode(&taskstatus)
	if err != nil {
		return nil, err
	}

	return &taskstatus, nil
}

func GetAllTaskStatus(ctx context.Context, limit, offset int) ([]models.TaskStatus, int64, error) {
	var taskstatus []models.TaskStatus

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

	cursor, err := database.TaskStatus.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Decode the users from the cursor
	if err = cursor.All(ctx, &taskstatus); err != nil {
		return nil, 0, err
	}

	count, err := database.TaskStatus.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return taskstatus, count, nil
}

func UpdateTaskStatus(ctx context.Context, id primitive.ObjectID, taskstatus models.TaskStatus) (*mongo.UpdateResult, error) {
	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	updateFields := bson.M{"UpdatedAt": time.Now()}

	updateFields["Name"] = taskstatus.Name
	updateFields["Description"] = taskstatus.Description

	update := bson.M{"$set": updateFields}
	result, err := database.TaskStatus.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteTaskStatus(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {
	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	result, err := database.TaskStatus.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}
