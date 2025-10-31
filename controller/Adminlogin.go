package controller

import (
	"context"
	"net/http"
	"time"
	"yadhronics-blog/database"
	"yadhronics-blog/response"
	"yadhronics-blog/security"
	"yadhronics-blog/service"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func AdminLogin(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	var adminlogin struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&adminlogin); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Failed to parse request body",
			ErrorTime:    time.Now(),
		})
	}

	//fetch data from DB
	err := service.AdminLogin(ctx, adminlogin.Email, adminlogin.Password)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// generate cookie
	cookie, err := security.GenerateJWTCookie(adminlogin.Email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// Return a success response
	return c.Status(http.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data:          cookie,
	})
}

func AdminCreatePassword(c *fiber.Ctx) error {
	var admincreatepassword struct {
		Password string `json:"password"`
	}

	if err := c.BodyParser(&admincreatepassword); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Failed to parse request body",
			ErrorTime:    time.Now(),
		})
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admincreatepassword.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	return c.Status(http.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data:          string(hashedPassword),
	})
}

func AdminValidate(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data:          nil,
	})
}
