package auth

import (
	"fiber-bbs/models/user"
	"fiber-bbs/requests"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type RegisterHandler struct{}

func (r *RegisterHandler) ShowRegistrationForm(c *fiber.Ctx) error {
	return c.Render("auth/register", fiber.Map{
		"Title": "注册",
	})
}
func (r RegisterHandler) Register(c *fiber.Ctx) error {
	user := user.User{}
	if err := c.BodyParser(&user); err != nil {
		fmt.Println(err)
	}
	errors := requests.ValidateRegistrationForm(&user, c)
	if len(errors) > 0 {
		return c.Render("auth/register", fiber.Map{
			"Title":  "注册",
			"Errors": errors,
			"User":   &user,
		})
	}
	user.Create()
	return c.Redirect("/")
}
