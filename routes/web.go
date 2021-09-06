package routes

import (
	"fiber-bbs/handlers"
	"fiber-bbs/handlers/auth"
	"image/color"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/steambap/captcha"
)

func RegisterWebRoutes(app *fiber.App) {
	home := &handlers.HomeHandler{}
	app.Get("/", home.Index)
	app.Get("/captcha", func(c *fiber.Ctx) error {

		store := session.New()
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		data, _ := captcha.New(100, 30, func(options *captcha.Options) {
			options.CharPreset = "1234567890"
			options.CurveNumber = 2
			options.TextLength = 4
			options.Palette = color.Palette{}
		})

		sess.Set("captcha", data.Text)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		return data.WriteImage(c.Response().BodyWriter())
	})
	register := &auth.RegisterHandler{}
	app.Get("/register", register.ShowRegistrationForm)
	app.Post("/register", register.Register)
	login := &auth.LoginHandler{}
	app.Get("/login", login.ShowLoginForm)
	app.Post("/login", login.Login)
}
