package handlers

import (
	"fiber-bbs/pkgs/auth"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
}

func (h *HomeHandler) Index(c *fiber.Ctx) error {
	user := auth.User(c)

	return c.Render("index", fiber.Map{
		"Title": "首页",
		"User":  user,
	})
}
