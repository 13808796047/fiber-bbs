package handlers

import (
	"fiber-bbs/models/topic"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/slava-vishnyakov/paginator"
	"github.com/spf13/cast"
	"html/template"
)

type TopicHandler struct {
}

func (t *TopicHandler) Index(c *fiber.Ctx) error {
	data := map[string]interface{}{
		"page":     c.Query("page", "1"),
		"per_page": c.Query("per_page", "10"),
	}
	topics, count, _ := topic.GetList(data)

	p := paginator.Paginator(cast.ToInt(c.Query("page", "1")), 10, cast.ToInt(count), 2)
	res := ""
	for _, page := range p {
		var class string
		var disable string
		if page.IsActive {
			class = "active"
		} else {
			class = ""
		}
		if page.IsDisabled {
			disable = "disabled"
		} else {
			disable = ""
		}

		res += fmt.Sprintf(`<li class="page-item %s %s"><a class="page-link" href="?page=%v">%s</a></li>`, class, disable, page.Page, page.Label)
	}

	return c.Render("topics/index", &fiber.Map{
		"Topics":     topics,
		"page_class": "topics-index-page",
		"Page":       template.HTML(res),
	})
}
