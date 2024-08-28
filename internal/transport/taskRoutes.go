package transport

import (
	"modules/internal/models"
	"modules/internal/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// создание таска
func insertTask(c *fiber.Ctx) error {
	var curUser models.Task

	if err := c.BodyParser(&curUser); err != nil {
		return models.ResponseTask{}.BadRequest().GetResponse(c)
	}

	return services.TaskInsert(curUser).GetResponse(c)
}

// получить все таски для просмотра
func getAllTasks(c *fiber.Ctx) error {
	var curUser models.Task
	return services.TaskShow(curUser).GetResponse(c)
}

// получить конкретный таск по айди
func getTask(c *fiber.Ctx) error {
	var curUser models.Task

	tempId := c.Params("id")
	idNew, err := strconv.Atoi(tempId)
	if err != nil {
		return models.ResponseTask{}.NotFound().GetResponse(c)
	}
	curUser.Id = idNew

	return services.TaskShow(curUser).GetResponse(c)
}

// изменить конкретный таск
func updateTask(c *fiber.Ctx) error {
	var curUser models.Task

	if err := c.BodyParser(&curUser); err != nil {
		return models.ResponseTask{}.BadRequest().GetResponse(c)
	}

	tempId := c.Params("id")
	idNew, err := strconv.Atoi(tempId)
	if err != nil {
		return models.ResponseTask{}.NotFound().GetResponse(c)
	}
	curUser.Id = idNew

	return services.TaskUpdate(curUser).GetResponse(c)
}

// удалить конкретный таск
func deleteTask(c *fiber.Ctx) error {
	var curUser models.Task

	tempId := c.Params("id")
	idNew, err := strconv.Atoi(tempId)
	if err != nil {
		return models.ResponseTask{}.NotFound().GetResponse(c)
	}
	curUser.Id = idNew

	return services.TaskDelete(curUser).GetResponse(c)
}
