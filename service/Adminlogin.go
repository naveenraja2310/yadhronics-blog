package service

import (
	"context"
	"yadhronics-blog/database"
	"yadhronics-blog/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(ctx context.Context, email string, password string) error {

	var admin models.Admin
	err := database.AdminLogin.FindOne(ctx, bson.M{"email": email}).Decode(&admin)
	if err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}
