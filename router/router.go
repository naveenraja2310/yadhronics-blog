package router

import (
	"yadhronics-blog/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GetRouter() *fiber.App {
	app := fiber.New()

	app.Use(logger.New())
	app.Use(recover.New())

	app.Post("/blog", controller.CreateBlog)
	app.Put("/blog/:id", controller.UpdateBlog)
	app.Get("/blog/:id", controller.GetBlogById)
	app.Get("/blog", controller.GetAllBlogs)
	app.Get("/blog-group", controller.GetBlogGroup)
	app.Delete("/blog/:id", controller.DeleteBlog)
	return app
}
