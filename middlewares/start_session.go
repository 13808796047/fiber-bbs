package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"log"
)

var Sess *session.Session
var err error

func StartSession() fiber.Handler {
	return func(c *fiber.Ctx) error {
		Sess, err = session.New().Get(c)
		if err != nil {
			log.Fatal(err)
		}
		return c.Next()
	}
}
