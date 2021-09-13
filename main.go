package main

import (
	"fiber-bbs/bootstrap"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"strings"
)

func main() {
	engine := html.New("./views", ".html")
	// Creating a template with function hasPermission
	engine.AddFunc("active_class", func(path, current_path string) bool {
		return strings.Contains(path, current_path)
	})
	//engine.AddFunc("if_route", helpers.IfRoute)
	//engine.AddFunc("test01", func() {
	//	var route fiber.Route
	//	route.Path
	//})
	//// SetRoute 设置路由实例,以供RouteName2URL等函数使用
	//func SetRoute(r *mux.Router) {
	//	route = r
	//}
	//
	//// RouteName2URL 通过路由名称来获取URL
	//func RouteName2URL(routeName string, pairs ...string) string {
	//
	//	url, err := route.Get(routeName).URL(pairs...)
	//	if err != nil {
	//	logger.LogError(err)
	//	return ""
	//}
	//
	//	return config.GetString("app.url") + url.String()
	//}
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/app",
	})

	app.Static("/static", "static")
	bootstrap.SetupRoute(app)
	bootstrap.SetupDB()
	app.Listen(":3000")
}
