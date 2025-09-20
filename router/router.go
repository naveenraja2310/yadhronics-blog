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
	app.Delete("/blog/:id", controller.DeleteBlog)
	// app.Get("/taskstatus/:id", controller.GetTaskStatusById)
	// app.Get("/taskstatus", controller.GetAllTaskStatus)
	// app.Put("/taskstatus/:id", controller.UpdateTaskStatus)
	// app.Delete("/taskstatus/:id", controller.DeleteTaskStatus)
	return app
}
