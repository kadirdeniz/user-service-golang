package tools

import (
	"sync"

	"github.com/gofiber/fiber/v2"
)

var mutex = &sync.Mutex{}

var app *fiber.App

func NewServer() *fiber.App {
	mutex.Lock()
	defer mutex.Unlock()

	if app == nil {
		app = fiber.New()
	}
	return app
}
