package handlers

import (
	"fiber-bbs/models/topic"
	"fiber-bbs/pkgs/pagination"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"html/template"
	"log"
)

type TopicHandler struct {
}

func (t *TopicHandler) Index(c *fiber.Ctx) error {

	log.Println(c.Path())
	page := cast.ToInt(c.Query("page", "1"))
	per_page := cast.ToInt(c.Query("per_page", "5"))
	order := c.Query("order")
	maps := make(map[string]interface{})
	topics, count, _ := topic.GetList(page, per_page, maps, order)
	res := pagination.Pagination(page, cast.ToInt(count), per_page)
	return c.Render("topics/index", &fiber.Map{
		"Topics":     topics,
		"page_class": "topics-index-page",
		"Page":       template.HTML(res),
		"Title":      "话题列表",
	})
}
