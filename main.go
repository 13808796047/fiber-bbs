package main

import (
	"fiber-bbs/bootstrap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"html/template"
	"strings"
)

func main() {
	engine := html.New("./views", ".html")
	// Creating a template with function hasPermission
	engine.AddFunc("urlContain", func(current_path, path string) bool {
		return strings.Contains(current_path, path)
	})
	engine.AddFunc("str2html", func(str string) template.HTML {
		return template.HTML(str)
	})
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/app",
	})

	app.Static("/static", "static")
	bootstrap.SetupRoute(app)
	bootstrap.SetupDB()
	app.Listen(":3000")
}
