package handlers

import (
	"fiber-bbs/database"
	"fiber-bbs/models/category"
	"fiber-bbs/models/topic"
	"fiber-bbs/models/user"
	"fiber-bbs/pkgs/auth"
	"fiber-bbs/pkgs/pagination"
	"fiber-bbs/pkgs/upload"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"html/template"
	"log"
)

type TopicHandler struct {
}

func (t *TopicHandler) Index(c *fiber.Ctx) error {
	page := cast.ToInt(c.Query("page", "1"))
	per_page := cast.ToInt(c.Query("per_page", "5"))
	order := c.Query("order")
	maps := make(map[string]interface{})
	topics, count, _ := topic.GetList(page, per_page, maps, order)
	res := pagination.Pagination(page, cast.ToInt(count), per_page)
	return c.Render("topics/index", &fiber.Map{
		"Topics":       topics,
		"page_class":   "topics-index-page",
		"Page":         template.HTML(res),
		"Title":        "话题列表",
		"current_path": c.OriginalURL(),
		"User":         auth.User(c),
	})
}
func (t *TopicHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id", "0")
	topic, err := topic.Get(id)
	if err != nil {
		log.Fatal(err)
	}
	return c.Render("topics/show", &fiber.Map{
		"topic":        &topic,
		"page_class":   "topics-show-page",
		"current_path": c.OriginalURL(),
	})
}
func (t *TopicHandler) Create(c *fiber.Ctx) error {
	categories, _ := category.List()
	return c.Render("topics/create_and_edit", &fiber.Map{
		"Categories":   &categories,
		"current_path": c.OriginalURL(),
	})
}
func (t *TopicHandler) Store(c *fiber.Ctx) error {
	category, err := category.Get(cast.ToInt(c.FormValue("category_id")))
	if err != nil {
		log.Fatal(err)
	}
	u, _ := user.Get("1")
	topic := topic.Topic{
		Category: category,
		User:     &u,
	}
	if err := c.BodyParser(&topic); err != nil {
		log.Fatal(err)
	}

	database.DB.Create(&topic)
	return c.Redirect("/topics/show/"+cast.ToString(topic.ID), 302)
}
func (t *TopicHandler) UploadImage(c *fiber.Ctx) error {
	h, err := c.FormFile("upload_file")
	if err != nil {
		return err
	}
	//user := auth.User(c)
	path, err := upload.Save(h, "topics", "1")
	data := map[string]interface{}{
		"success":   false,
		"msg":       "上传失败",
		"file_path": "",
	}

	if err != nil {
		data["msg"] = err
		return c.JSON(data)
	}
	err = c.SaveFile(h, path)
	if err != nil {
		data["msg"] = err
		return c.JSON(data)
	}
	data = map[string]interface{}{
		"success":   true,
		"msg":       "上传成功",
		"file_path": path,
	}
	return c.JSON(data)
}
