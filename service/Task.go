package service

import (
	"context"
	"errors"
	"time"

	"task_management/database"
	"task_management/models"
	"task_management/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateTask(ctx context.Context, task models.Task) (*mongo.InsertOneResult, error) {
	task.CreatedAt = time.Now()

	result, err := database.Tasks.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetTaskByID(ctx context.Context, id primitive.ObjectID) (models.Task, error) {
	var task models.Task

	if !utils.CheckIfExistsByID(ctx, database.Tasks, id) {
		return models.Task{}, errors.New("the given id is invalid")
	}

	// Define the MongoDB aggregation pipeline
	pipeline := mongo.Pipeline{
		// Match the specific task by its ID
		{{Key: "$match", Value: bson.D{{Key: "_id", Value: id}}}},
		// Join with the 'users' collection to fetch the FromUser details
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "fromUserId"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "fromUser"},
		}}},
		// Join with the 'users' collection to fetch the ToUser details
		{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "users"},
			{Key: "localField", Value: "toUserId"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "toUser"},
		}}},
		// Unwind the fromUser and toUser arrays (assuming there's one user per task)
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$fromUser"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},
		{{Key: "$unwind", Value: bson.D{{Key: "path", Value: "$toUser"}, {Key: "preserveNullAndEmptyArrays", Value: true}}}},
	}

	// Execute the aggregation
	cursor, err := database.Tasks.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return models.Task{}, err
	}

	// Decode the result into a Task struct
	if cursor.Next(context.TODO()) {
		if err := cursor.Decode(&task); err != nil {
			return models.Task{}, err
		}
	}
	return task, nil
}
