package routes

import (
	"fiber-bbs/handlers"
	"fiber-bbs/handlers/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/steambap/captcha"
	"io"
)

type captchare struct {
	w *io.Writer
}

func RegisterWebRoutes(app *fiber.App)   {
	home := &handlers.HomeHandler{}
	app.Get("/",home.Index)
	app.Get("/captcha", func(c *fiber.Ctx) error {
		type w *io.Writer
		store := session.New()
		sess, err := store.Get(c)
		if err != nil {
			panic(err)
		}

		data, _ := captcha.New(150, 50)

		sess.Set("captcha", data.Text)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		return data.WriteImage(captchare{w: io.Writer()})
	})
	register := &auth.RegisterHandler{}
	app.Get("/register",register.ShowRegistrationForm)
	app.Post("/register",register.Register)
	login := &auth.LoginHandler{}
	app.Get("/login",login.ShowLoginForm)
	app.Post("/login",login.Login)
}