package handlers

import (
	"fiber-bbs/models/category"
	"fiber-bbs/models/topic"
	"fiber-bbs/pkgs/pagination"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"html/template"
	"log"
)

type CategoryHandler struct {
}

func (cate *CategoryHandler) Show(c *fiber.Ctx) error {
	category_id := c.Params("id")
	page := cast.ToInt(c.Query("page", "1"))
	per_page := cast.ToInt(c.Query("per_page", "5"))
	order := c.Query("order")
	result, err := category.Get(cast.ToInt(category_id))
	if err != nil {
		log.Println(err)
	}
	maps := map[string]interface{}{
		"category_id": result.ID,
	}
	topics, count, err := topic.GetList(page, per_page, maps, order)
	if err != nil {
		log.Println(err)
	}

	res := pagination.Pagination(page, cast.ToInt(count), per_page)

	return c.Render("topics/index", &fiber.Map{
		"Topics":       &topics,
		"page_class":   "topics-index-page",
		"Page":         template.HTML(res),
		"Title":        result.Name,
		"Category":     result,
		"current_path": c.OriginalURL(),
	})
}
