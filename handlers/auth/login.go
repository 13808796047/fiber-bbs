package auth

import "github.com/gofiber/fiber/v2"

type LoginHandler struct{}

func (this *LoginHandler) ShowLoginForm(c *fiber.Ctx) error {
	return c.Render("auth/login", fiber.Map{
		"Title": "登录",
	})
}
func (this *LoginHandler) Login(c *fiber.Ctx) error {

	return c.Redirect("/")
}
