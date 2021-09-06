package auth

import (
	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct {}

func (r *RegisterHandler)ShowRegistrationForm(c *fiber.Ctx) error {


	return c.Render("auth/register",fiber.Map{
		"Title": "注册",
	})
}
func (r RegisterHandler) Register(c *fiber.Ctx) error {
	return c.Redirect("/")
}