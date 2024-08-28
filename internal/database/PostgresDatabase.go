package database

import (
	"fmt"
	"modules/internal/config"
	"modules/internal/models"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// запуск базы данных. сначала проверка есть ли нужная бд и её создание, потом миграция, и добавляем экземпляр базы в глобальную переменную.
func (currentlDB *PostgresDatabase) Run(config config.MainConfig) {
	currentlDB.OpenConnection(config.PostgresDB)
	currentlDB.StartMigration()
	currentlDB.GlobalSet()
}

// миграция таблицы
func (currentlDB *PostgresDatabase) StartMigration() {
	currentlDB.Instance.AutoMigrate(models.Task{})
	log.Debug("migration complete")
}

// открыть соединение с базой
func (currentlDB *PostgresDatabase) OpenConnection(config config.PostgresConfig) {

	err := currentlDB.checkDatabaseCreated(config)
	if err != nil {
		panic("not connection to db")
	}

	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, config.Name)

	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		panic("not connection to db")
	}

	currentlDB.Instance = db
}

// проверка если есть база данных. если нет, то создать.
func (currentlDB *PostgresDatabase) checkDatabaseCreated(config config.PostgresConfig) error {

	// открытие соеднение с базой по стандарту
	connectStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", config.User, config.Pass, config.Host, config.Port, "postgres")
	db, err := gorm.Open(postgres.Open(connectStr), &gorm.Config{})
	if err != nil {
		log.Error("database don't open")
		return models.ResponseGlobal{}.InternalError()
	}

	// закрытие бд
	sql, _ := db.DB()
	defer func() {
		_ = sql.Close()
	}()

	// проверка если есть нужная нам база данных
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", config.Name)
	rs := db.Raw(stmt)
	if rs.Error != nil {
		log.Error("error, dont read bd")
		return models.ResponseGlobal{}.InternalError()
	}

	// если нет, то создать
	var rec = make(map[string]interface{})
	if rs.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", config.Name)
		if rs := db.Exec(stmt); rs.Error != nil {
			log.Error("error, dont create a database")
			return models.ResponseGlobal{}.InternalError()
		}
	}

	return nil
}

// текущий экземпляр таблицы в глобальную переменную
func (currentlDB *PostgresDatabase) GlobalSet() {
	GlobalPostgres = currentlDB
}

// парсинг интерфейса Table в таблицу Task
func (currentlDB *PostgresDatabase) getData(temp models.Table) (models.Task, models.Response) {
	person, ok := temp.(*models.Task)
	if ok == false {
		return models.Task{}, models.ResponseTask{}.InternalError()
	}
	return *person, nil
}
