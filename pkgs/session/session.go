package session

import (
	"github.com/gofiber/fiber/v2/middleware/session"
)

var SessionStore = session.New()

// // 存入session
// func SessionStart(c *fiber.Ctx) *session.Session {
// 	currSession, err := SessionStore.Get(c)
// 	defer currSession.Save()
// 	if err != nil {
// 		log.Panicln("currSession获取错误：", err)
// 	}
// 	err = currSession.Regenerate()
// 	if err != nil {
// 		log.Panicln("currSession获取错误：", err)
// 	}
// 	return currSession
// }
