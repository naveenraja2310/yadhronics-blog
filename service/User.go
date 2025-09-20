package service

import (
	"context"
	"errors"
	"fmt"
	"time"
	"yadhronics-blog/database"
	"yadhronics-blog/models"
	"yadhronics-blog/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	user.CreatedAt = time.Now()

	query := bson.M{
		"$or": []bson.M{
			{"Email": user.Email},
			{"PhoneNumber": user.PhoneNumber},
		},
	}

	if utils.IsDuplicate(ctx, database.User, query) {
		return nil, fmt.Errorf("email / phonenumber is already available")
	}

	result, err := database.User.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetUserByID(ctx context.Context, id primitive.ObjectID) (*models.User, error) {
	var user models.User

	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	err := database.User.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUser(ctx context.Context, limit, offset int) ([]models.User, int64, error) {
	var user []models.User

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

	cursor, err := database.User.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	// Decode the users from the cursor
	if err = cursor.All(ctx, &user); err != nil {
		return nil, 0, err
	}

	count, err := database.User.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, err
	}

	return user, count, nil
}

func UpdateUser(ctx context.Context, id primitive.ObjectID, user models.User) (*mongo.UpdateResult, error) {

	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	updateFields := bson.M{"UpdatedAt": time.Now()}

	updateFields["Name"] = user.Name
	updateFields["Email"] = user.Email
	updateFields["PhoneNumber"] = user.PhoneNumber
	updateFields["IsVerified"] = user.IsVerified
	updateFields["IsAdmin"] = user.IsAdmin
	updateFields["Photo"] = user.Photo

	update := bson.M{"$set": updateFields}
	result, err := database.User.UpdateOne(ctx, bson.M{"_id": id}, update)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteUser(ctx context.Context, id primitive.ObjectID) (*mongo.DeleteResult, error) {

	if !utils.CheckIfExistsByID(ctx, database.TaskStatus, id) {
		return nil, errors.New("the given id is invalid")
	}

	result, err := database.User.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return nil, err
	}

	return result, nil
}
