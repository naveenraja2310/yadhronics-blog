package database

import (
	"context"
	"yadhronics-blog/models"
	"yadhronics-blog/settings"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoClient       *mongo.Client
	Blogs, AdminLogin *mongo.Collection
	ContextTime       int = 5
)

func InitDB(config settings.Configuration) error {
	clientOptions := options.Client().ApplyURI(config.DBURI)

	// Connect to MongoDB using the client options
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return err
	}

	blogs := models.Blogs{}
	Blogs = client.Database(config.DB_NAME).Collection(blogs.TableName())

	adminlogin := models.Admin{}
	AdminLogin = client.Database(config.DB_NAME).Collection(adminlogin.TableName())

	return nil
}
