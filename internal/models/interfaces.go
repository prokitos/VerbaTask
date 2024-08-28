package models

import (
	"modules/internal/config"

	"github.com/gofiber/fiber/v2"
)

// интерфейс для респонсов
type Response interface {
	GetResponse(c *fiber.Ctx) error
	Validate() bool
}

// интерфейс для таблиц
type Table interface {
	RecordCreate(GormDatabase) Response
	RecordDelete(GormDatabase) Response
	RecordShow(GormDatabase) Response
	RecordUpdate(GormDatabase) Response
}

// интерфейс для таблиц горма (постгрес)
type GormDatabase interface {
	OpenConnection(config.PostgresConfig)
	StartMigration()
	GlobalSet()
	CreateData(Table) Response
	DeleteData(Table) Response
	UpdateData(Table) Response
	ShowData(Table) Response
}
