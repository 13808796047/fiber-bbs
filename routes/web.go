package routes

import (
	"fiber-bbs/handlers"
	"fiber-bbs/handlers/auth"
	"image/color"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/steambap/captcha"
)

func RegisterWebRoutes(app *fiber.App) {
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))
	//app.Use(middlewares.StartSession())
	home := &handlers.HomeHandler{}
	app.Get("/", home.Index)

	app.Get("/captcha", func(c *fiber.Ctx) error {

		data, _ := captcha.New(100, 30, func(options *captcha.Options) {
			options.CharPreset = "1234567890"
			options.CurveNumber = 2
			options.TextLength = 4
			options.Palette = color.Palette{}
		})

		c.Cookie(&fiber.Cookie{
			Name:  "captcha",
			Value: data.Text,
		})
		return data.WriteImage(c.Response().BodyWriter())
	})
	register := &auth.RegisterHandler{}
	app.Get("/register", register.ShowRegistrationForm)
	app.Post("/register", register.Register)
	login := &auth.LoginHandler{}
	app.Get("/login", login.ShowLoginForm)
	app.Post("/login", login.Login)
	app.Post("/logout", login.Logout)
	topic := &handlers.TopicHandler{}
	app.Get("/topics", topic.Index)
}
