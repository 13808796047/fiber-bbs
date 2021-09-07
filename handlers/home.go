package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	*Base
}

func (h *HomeHandler) Index(c *fiber.Ctx) error {
	user := h.User(c)

	return c.Render("index", fiber.Map{
		"Title": "首页",
		"User":  user,
	})
}
