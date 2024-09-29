package controller

import (
	"context"
	"net/http"
	"task_management/database"
	"task_management/models"
	"task_management/response"
	"task_management/service"
	"task_management/utils"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateTask(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	//parsing a request body
	var task models.Task
	if err := c.BodyParser(&task); err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: "Failed to parse request body",
			ErrorTime:    time.Now(),
		})
	}

	//saving data in db
	result, err := service.CreateTask(ctx, task)
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

func GetTaskById(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	idParam := utils.StringToObjectID(c.Params("id"))

	//fetch data from DB
	result, err := service.GetTaskByID(ctx, idParam)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// Return a success response
	return c.Status(http.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data:          &fiber.Map{"data": result},
	})
}
