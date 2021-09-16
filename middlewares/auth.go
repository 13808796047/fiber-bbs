package middlewares

import (
	"fiber-bbs/pkgs/auth"
	"github.com/gofiber/fiber/v2"
)

func CheckAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		if !auth.Check(c) {
			return c.Redirect("/login")
		}
		return c.Next()
	}
}
