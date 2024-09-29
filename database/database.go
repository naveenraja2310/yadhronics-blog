package database

import (
	"context"
	"fmt"
	"task_management/models"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient                       *mongo.Client
	TaskStatus, TaskType, User, Tasks *mongo.Collection
	ContextTime                       int = 5
)

func InitDB() error {
	clientOptions := options.Client().ApplyURI("mongodb+srv://naveenraja2310:yx79hpUr6DOBJ70M@cluster0.ria4e.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")

	// Connect to MongoDB using the client options
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		// Log and terminate the program if unable to connect to MongoDB
		return err
	}
	fmt.Println("client", client)
	task_status := models.TaskStatus{}
	TaskStatus = client.Database("Task_Manegement").Collection(task_status.TableName())

	task_type := models.TaskType{}
	TaskType = client.Database("Task_Manegement").Collection(task_type.TableName())

	user := models.User{}
	User = client.Database("Task_Manegement").Collection(user.TableName())

	task := models.Task{}
	Tasks = client.Database("Task_Manegement").Collection(task.TableName())

	return nil
}
