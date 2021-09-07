package auth

import (
	"errors"
	"fiber-bbs/pkgs/auth"

	"github.com/gofiber/fiber/v2"
)

type LoginHandler struct {
}

func (l *LoginHandler) ShowLoginForm(c *fiber.Ctx) error {

	return c.Render("auth/login", fiber.Map{
		"Title": "登录",
	})
}
func (l *LoginHandler) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")
	if err := auth.Attempt(c, email, password); err == nil {
		return c.Redirect("/")
	}
	return c.Render("auth/login", fiber.Map{
		"Title":    "登录",
		"Errors":   errors.New("账号或密码错误"),
		"Email":    email,
		"Password": password,
	})
}
func (l *LoginHandler) Logout(c *fiber.Ctx) error {
	auth.Logout(c)
	return c.Redirect("/")
}
