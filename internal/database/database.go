package database

import (
	"gorm.io/gorm"
)

// глобальное соединение с базой данных.
var GlobalPostgres *PostgresDatabase

type PostgresDatabase struct {
	Instance *gorm.DB
}
