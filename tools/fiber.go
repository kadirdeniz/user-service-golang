package tools

import (
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

var mutex = &sync.Mutex{}

var app *fiber.App

func NewServer() *fiber.App {
	mutex.Lock()
	defer mutex.Unlock()

	if app == nil {
		app = fiber.New()
		app.Static("/swagger/doc.json", "./docs/swagger.json")

		app.Get("/swagger/*", swagger.HandlerDefault)
	}
	return app
}
