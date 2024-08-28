package database

import (
	"fmt"
	"modules/internal/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

// тут DAO для таблицы Task (удаление, обновление, просмотр и добавление)

func (currentlDB *PostgresDatabase) CreateData(data models.Table) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	result := currentlDB.Instance.Clauses(clause.Returning{}).Create(&task)
	if result.Error != nil {
		log.Debug("create record error!")
		return models.ResponseTask{}.InternalError()
	}

	var finded []models.Task
	finded = append(finded, task)
	return models.ResponseTask{}.GoodCreate(finded)
}

func (currentlDB *PostgresDatabase) ShowData(data models.Table) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}
	var finded []models.Task

	results := currentlDB.Instance.Find(&finded, task)
	if results.Error != nil {
		log.Debug("show record error!")
		return models.ResponseTask{}.InternalError()
	}

	if len(finded) == 0 {
		log.Debug("nothing to show")
		return models.ResponseTask{}.NotFound()
	}

	return models.ResponseTask{}.GoodShow(finded)

}

func (currentlDB *PostgresDatabase) UpdateData(data models.Table) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	fmt.Println(task)

	result := currentlDB.Instance.Where("id = ?", task.Id).Clauses(clause.Returning{}).Updates(&task)
	if result.Error != nil {
		log.Debug("update record error!")
		return models.ResponseTask{}.InternalError()
	}

	var finded []models.Task
	finded = append(finded, task)

	return models.ResponseTask{}.GoodUpdate(finded)
}

func (currentlDB *PostgresDatabase) DeleteData(data models.Table) models.Response {

	task, resp := currentlDB.getData(data)
	if resp != nil {
		return resp
	}

	result := currentlDB.Instance.Delete(&task, task.Id)
	if result.RowsAffected == 0 || result.Error != nil {
		return models.ResponseTask{}.NotFound()
	}

	return models.ResponseTask{}.GoodDelete()
}
