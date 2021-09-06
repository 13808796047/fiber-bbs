package handlers

import "github.com/gofiber/fiber/v2"

type HomeHandler struct {}
func (h *HomeHandler)Index(c *fiber.Ctx) error {
	return c.Render("index",fiber.Map{
		"Title": "首页",
	})
}