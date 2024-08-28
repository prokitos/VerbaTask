package services

import (
	"modules/internal/database"
	"modules/internal/models"
)

// конвертация времени в RFC3339, а затем создание нового таска
func TaskInsert(instance models.Task) models.Response {

	var worker models.TimeWorker
	curTime := worker.GetNowRfc()
	curDue := worker.ConvertStringToRfc(instance.DueDate)

	instance.CreatedAt = curTime
	instance.UpdatedAt = curTime
	instance.DueDate = curDue

	return instance.RecordCreate(database.GlobalPostgres)
}

// показать таски по указанным параметрам
func TaskShow(instance models.Task) models.Response {
	return instance.RecordShow(database.GlobalPostgres)
}

// конвертация времени в RFC3339, а затем обновить таск с нужным айди
func TaskUpdate(instance models.Task) models.Response {

	var worker models.TimeWorker
	curTime := worker.GetNowRfc()
	curDue := worker.ConvertStringToRfc(instance.DueDate)

	instance.UpdatedAt = curTime
	instance.DueDate = curDue
	instance.CreatedAt = ""

	return instance.RecordUpdate(database.GlobalPostgres)
}

// удалить таск с нужным айди
func TaskDelete(instance models.Task) models.Response {
	return instance.RecordDelete(database.GlobalPostgres)
}
