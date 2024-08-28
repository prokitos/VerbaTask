package models

import (
	"errors"
	"strconv"
)

// структура для стандартных ответов на ошибки (например при запуске базы данных)
type ResponseGlobal struct {
	Description string `json:"description"       		 example:"description"`
	Code        int    `json:"code"               		 example:"status"`
}

func (instance ResponseGlobal) BadRequest() error {
	instance.Code = 400
	instance.Description = "Bad Request"
	return errors.New(instance.getResult())
}
func (instance ResponseGlobal) InternalError() error {
	instance.Code = 500
	instance.Description = "Internal Server Error"
	return errors.New(instance.getResult())
}
func (instance ResponseGlobal) NotFound() error {
	instance.Code = 404
	instance.Description = "Not Found"
	return errors.New(instance.getResult())
}
func (instance ResponseGlobal) getResult() string {
	result := `{"code": ` + strconv.Itoa(instance.Code) + `, "description": "` + instance.Description + `"}`
	return result
}
