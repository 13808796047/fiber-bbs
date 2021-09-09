package helpers

import "github.com/gofiber/fiber/v2"

var r *fiber.Ctx

func ActiveClass(condition bool) string {

	var activeClass string
	if condition {
		activeClass = "active"
	} else {
		activeClass = ""
	}
	return activeClass
}
func IfRoute(path string) bool {
	return path == r.Path()
}
func IfRouteParam(param string) bool {
	return param == r.Params(param)
}
func CategoryNavActive(category_id string) string {
	return ActiveClass(IfRoute("categories/show") && IfRouteParam(category_id))
}
