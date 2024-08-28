package transport

import (
	"github.com/gofiber/fiber/v2"
)

// хэндлеры сервера
func SetHandlers(instance *fiber.App) {
	instance.Post("/tasks", insertTask)

	instance.Get("/tasks", getAllTasks)
	instance.Get("/tasks/:id", getTask)

	instance.Delete("/tasks/:id", deleteTask)
	instance.Put("/tasks/:id", updateTask)

}
