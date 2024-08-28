Проект билдится командой:   go build cmd/main.go 
ИЛИ запускается командой:   go run cmd/main.go
В папке config есть файл local.yaml, там можно поменять параметры сервера или базы данных.
после этого в папке проекта запускаем:  main.go 

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
Get:  	 http://localhost:8001/tasks/error

Delete:  http://localhost:8001/tasks/12

Put:	 http://localhost:8001/tasks/10
{
    "title": "complete",
    "description": "i already do something",
    "due_date":"2021-03-07"
}




вместо golint использовал revive.
так как он устарел и убран из большинства инструментов для анализа кода, а сам репозиторий golint остался только как зеркало.
