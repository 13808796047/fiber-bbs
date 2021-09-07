package handlers

import "github.com/gofiber/fiber/v2"

type TopicHandler struct {
}

func (t *TopicHandler) Index(c *fiber.Ctx) error {
	return c.Render("topic/index", &fiber.Map{})
}
