Проект билдится командой:   go build cmd/main.go 
ИЛИ запускается командой:   go run cmd/main.go

В папке config есть файл local.yaml, там можно поменять параметры сервера или базы данных.
при запуске создаётся БД с названием verba и таблица tasks
По стандарту установлен 8001 порт у сервера.

Роуты, и примеры запросов к ним:

Post:   http://localhost:8001/tasks 
{
    "title": "urgent",
    "description": "do something",
    "due_date":"2020.03.03"
}

Get:   	 http://localhost:8001/tasks

Get:  	 http://localhost:8001/tasks/5

Delete:  http://localhost:8001/tasks/12

Put:	 http://localhost:8001/tasks/10
{
    "title": "complete",
    "description": "i already do something",
    "due_date":"2021-03-07"
}


