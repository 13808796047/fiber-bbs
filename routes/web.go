package routes

import (
	"fiber-bbs/handlers"
	"fiber-bbs/handlers/auth"
	"fiber-bbs/middlewares"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/steambap/captcha"
	"image/color"
	"time"
)

// Timer will measure how long it takes before a response is returned
func Timer() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// start timer
		start := time.Now()
		// next routes
		err := c.Next()
		// stop timer
		stop := time.Now()
		// Do something with response
		c.Append("Server-Timing", fmt.Sprintf("app;dur=%v", stop.Sub(start).String()))
		// return stack error if exist
		return err
	}
}
func RegisterWebRoutes(app *fiber.App) {
	//app.Use(compress.New())
	//app.Use(logger.New())
	//app.Use(cache.New())
	app.Use(encryptcookie.New(encryptcookie.Config{
		Key: encryptcookie.GenerateKey(),
	}))
	//app.Use(Timer())
	//app.Use(cache.New(cache.Config{
	//	Next: func(c *fiber.Ctx) bool {
	//		return c.Query("refresh") == "true"
	//	},
	//	Expiration:   30 * time.Minute,
	//	CacheControl: true,
	//}))
	home := &handlers.HomeHandler{}
	app.Get("/", home.Index)
	//app.Use(Timer())
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

	topics := app.Group("/topics", middlewares.CheckAuth())
	{
		topic := &handlers.TopicHandler{}
		topics.Get("/", topic.Index)
		topics.Get("/show/:id", topic.Show)
		topics.Get("/create", topic.Create)
		topics.Post("/", topic.Store)
		topics.Post("/upload_image", topic.UploadImage)
	}

	categories := app.Group("/categories", middlewares.CheckAuth())
	{
		category := &handlers.CategoryHandler{}
		categories.Get("/show/:id", category.Show)
	}
}
