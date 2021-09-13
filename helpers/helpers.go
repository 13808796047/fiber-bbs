package helpers

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

var r *fiber.Route

func Initiail(route *fiber.Route) {
	r = route
}
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

	log.Println(r.Path)
	return path == r.Path

}

//func IfRouteParam(param string) bool {
//	return param == r.Params
//}
//func CategoryNavActive(category_id string) string {
//	return ActiveClass(IfRoute("categories/show") && IfRouteParam(category_id))
//}
