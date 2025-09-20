package controller

import (
	"context"
	"net/http"
	"strconv"
	"time"
	"yadhronics-blog/database"
	"yadhronics-blog/models"
	"yadhronics-blog/response"
	"yadhronics-blog/service"
	"yadhronics-blog/utils"

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
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// Return a success response with the created objectid
	return c.Status(http.StatusCreated).JSON(response.SuccessResponse{
		StatusCode:    http.StatusCreated,
		StatusMessage: "success",
		Data:          result,
	})
}

func UpdateBlog(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	idParam := utils.StringToObjectID(c.Params("id"))

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
	result, err := service.UpdateBlog(ctx, blog, idParam)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusInternalServerError,
			ErrorMessage: err.Error(),
			ErrorTime:    time.Now(),
		})
	}

	// Return a success response with the created objectid
	return c.Status(http.StatusOK).JSON(response.SuccessResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: "success",
		Data:          result,
	})
}

func GetBlogById(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	idParam := utils.StringToObjectID(c.Params("id"))

	//fetch data from DB
	result, err := service.GetBlogByID(ctx, idParam)
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
		Data:          result,
	})
}

func DeleteBlog(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	idParam := utils.StringToObjectID(c.Params("id"))

	//fetch data from DB
	err := service.DeleteBlog(ctx, idParam)
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
		Data:          nil,
	})
}

func GetAllBlogs(c *fiber.Ctx) error {
	//creating a context
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	limit, limiterr := strconv.Atoi(c.Query("limit"))
	if limiterr != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: limiterr.Error(),
			ErrorTime:    time.Now(),
		})
	}

	pagenumber, pagenumbererr := strconv.Atoi(c.Query("page"))
	if pagenumbererr != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: pagenumbererr.Error(),
			ErrorTime:    time.Now(),
		})
	}

	search := c.Query("search")
	category := c.Query("category")
	status := c.Query("status")

	offset := (pagenumber - 1) * limit

	//fetch data from DB
	result, count, err := service.GetAllBlogs(ctx, int64(limit), int64(offset), search, category, status)
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
		Data:          &fiber.Map{"blogs": result, "total_count": count},
	})
}

func GetBlogGroup(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(database.ContextTime)*time.Second)
	defer cancel()

	limit, limiterr := strconv.Atoi(c.Query("limit"))
	if limiterr != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: limiterr.Error(),
			ErrorTime:    time.Now(),
		})
	}

	pagenumber, pagenumbererr := strconv.Atoi(c.Query("page"))
	if pagenumbererr != nil {
		return c.Status(http.StatusBadRequest).JSON(response.ErrorResponse{
			ApiPath:      c.OriginalURL(),
			ErrorCode:    http.StatusBadRequest,
			ErrorMessage: pagenumbererr.Error(),
			ErrorTime:    time.Now(),
		})
	}

	offset := (pagenumber - 1) * limit

	result, err := service.GetBlogGroup(ctx, int64(limit), int64(offset))
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
		Data:          result,
	})
}
