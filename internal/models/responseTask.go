package models

import "github.com/gofiber/fiber/v2"

// структура для ответов на ошибки у таблицы Task
type ResponseTask struct {
	Description string `json:"description"        example:"description"`
	Code        int    `json:"code"               example:"status"`
	Users       []Task `json:"task,omitempty"    example:"...."`
}

func (instance ResponseTask) GoodCreate(curTasks []Task) Response {
	instance.Code = 201
	instance.Description = "Created"
	instance.Users = curTasks
	return instance.convertResponse()
}
func (instance ResponseTask) GoodShow(curUser []Task) Response {
	instance.Code = 200
	instance.Description = "Ok"
	instance.Users = curUser
	return instance.convertResponse()
}
func (instance ResponseTask) GoodDelete() Response {
	instance.Code = 204
	instance.Description = "No Content"
	instance.Users = nil
	return instance.convertResponse()
}
func (instance ResponseTask) GoodUpdate(curUser []Task) Response {
	instance.Code = 200
	instance.Description = "Ok"
	instance.Users = curUser
	return instance.convertResponse()
}

func (instance ResponseTask) InternalError() Response {
	instance.Code = 500
	instance.Description = "Internal Server Error"
	instance.Users = nil
	return instance.convertResponse()
}
func (instance ResponseTask) NotFound() Response {
	instance.Code = 404
	instance.Description = "Not Found"
	instance.Users = nil
	return instance.convertResponse()
}
func (instance ResponseTask) BadRequest() Response {
	instance.Code = 400
	instance.Description = "Bad Request"
	instance.Users = nil
	return instance.convertResponse()
}

func (instance ResponseTask) GetResponse(c *fiber.Ctx) error {
	return c.Status(instance.Code).JSON(instance)
}

func (instance ResponseTask) Validate() bool {
	if instance.Code >= 200 && instance.Code <= 300 {
		return true
	}
	return false
}

func (instance ResponseTask) convertResponse() Response {
	var temp Response = instance
	return temp
}
