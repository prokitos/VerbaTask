package database

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// глобальное соединение с базой данных.
var GlobalPostgres *PostgresDatabase

type PostgresDatabase struct {
	Instance *gorm.DB
}

func checkError(err error) {
	if err != nil {
		log.Error(err)
	}
}
