package router

import (
	"task_management/controller"

	"github.com/gofiber/fiber/v2"
)

func GetRouter() *fiber.App {
	app := fiber.New()

	//Task Status
	app.Post("/taskstatus", controller.CreateTaskStatus)
	app.Get("/taskstatus/:id", controller.GetTaskStatusById)
	app.Get("/taskstatus", controller.GetAllTaskStatus)
	app.Put("/taskstatus/:id", controller.UpdateTaskStatus)
	app.Delete("/taskstatus/:id", controller.DeleteTaskStatus)

	//Task Type
	app.Post("/tasktype", controller.CreateTaskType)
	app.Get("/tasktype/:id", controller.GetTaskTypeById)
	app.Get("/tasktype", controller.GetAllTaskType)
	app.Put("/tasktype/:id", controller.UpdateTaskType)
	app.Delete("/tasktype/:id", controller.DeleteTaskType)

	//User
	app.Post("/user", controller.CreateUser)
	app.Get("/user/:id", controller.GetUserById)
	app.Get("/user", controller.GetAllUser)
	app.Put("/user/:id", controller.UpdateUser)
	app.Delete("/user/:id", controller.DeleteUser)

	return app
}
