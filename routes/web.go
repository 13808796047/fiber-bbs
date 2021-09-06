package routes

import (
	"fiber-bbs/handlers"
	"github.com/gofiber/fiber/v2"
)

func RegisterWebRoutes(app *fiber.App)   {
	home := &handlers.HomeHandler{}
	app.Get("/",home.Index)
}