package models

// структура Task. именно так она хранится в таблице.
type Task struct {
	ID          int    `json:"id" example:"12" gorm:"unique;primaryKey;autoIncrement"`
	Title       string `json:"title" example:"today"`
	Description string `json:"description" example:"do something"`
	DueDate     string `json:"due_date" example:"11.11.2011"`
	CreatedAt   string `json:"created_at" example:"11.11.2011"`
	UpdatedAt   string `json:"updated_at" example:"11.11.2011"`
}

// Методы для вызова CRUD у базы данных.
func (instance *Task) RecordCreate(db GormDatabase) Response {
	return db.CreateData(instance)
}
func (instance *Task) RecordShow(db GormDatabase) Response {
	return db.ShowData(instance)
}
func (instance *Task) RecordDelete(db GormDatabase) Response {
	return db.DeleteData(instance)
}
func (instance *Task) RecordUpdate(db GormDatabase) Response {
	return db.UpdateData(instance)
}
