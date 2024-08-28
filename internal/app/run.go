package app

import (
	"fmt"
	"log"
	"modules/internal/transport"
	"time"

	"github.com/gofiber/fiber/v2"
)

type App struct {
	Server *fiber.App
}

// старт сервера. подключение хендлеров и настройка порта.
func (a *App) NewServer(port string) {
	app := fiber.New()
	a.Server = app
	a.setHandler()
	a.launchServer(port)
}

// выключение сервера при сбоях
func (a *App) Stop() {
	fmt.Println("Gracefully shutting down...")
	err := a.Server.ShutdownWithTimeout(20 * time.Second)
	if err != nil {
		panic("error at close server")
	}
}

func (a *App) setHandler() {
	transport.SetHandlers(a.Server)
}

func (a *App) launchServer(port string) {
	log.Fatal(a.Server.Listen(port))
}
