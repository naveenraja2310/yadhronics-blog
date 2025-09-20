package controller

import (
	"context"
	"net/http"
	"time"
	"yadhronics-blog/database"
	"yadhronics-blog/models"
	"yadhronics-blog/response"
	"yadhronics-blog/service"

	"github.com/gofiber/fiber/v2"
)

func CreateBlog(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	//parsing a request body
	var blog models.Blogs
	if err := c.BodyParser(&blog); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Failed to parse request body",
			ErrorTime:    time.Now(),
		})
	}

	//saving data in db
	result, err := service.CreateBlog(ctx, blog)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// Return a success response with the created objectid
	return c.Status(http.StatusCreated).JSON(response.SuccessResponse{
		StatusCode:    http.StatusCreated,
		StatusMessage: "success",
		Data:          &fiber.Map{"data": result},
	})
}
