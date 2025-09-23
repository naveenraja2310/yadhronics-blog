package router

import (
	"yadhronics-blog/controller"
	"yadhronics-blog/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GetRouter() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New(cors.Config{
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	app.Post("/blog", middleware.JWTMiddleware(), controller.CreateBlog)
	app.Put("/blog/:id", middleware.JWTMiddleware(), controller.UpdateBlog)
	app.Delete("/blog/:id", middleware.JWTMiddleware(), controller.DeleteBlog)

	app.Get("/blog/:id", controller.GetBlogById)
	app.Get("/blog", controller.GetAllBlogs)
	app.Get("/blog-group", controller.GetBlogGroup)

	app.Post("/adminlogin", controller.AdminLogin)
	app.Get("/adminvalidate", middleware.JWTMiddleware(), controller.AdminValidate)
	app.Get("/aws/presigned-url", middleware.JWTMiddleware(), controller.AwsPresignedURL)

	return app
}
