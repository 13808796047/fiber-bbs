package handlers

import (
	"errors"
	"fiber-bbs/models/user"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"gorm.io/gorm"
	"log"
)

type Base struct {
}

var SessionStore = session.New()

func (b *Base) _getUID(c *fiber.Ctx) interface{} {
	currSession, err := SessionStore.Get(c)
	if err != nil {
		log.Println(err)
	}
	_uid := currSession.Get("uid")

	err = currSession.Save()
	if err != nil {
		log.Println(err)
	}
	if _uid != nil {
		return _uid
	}
	return ""
}

// 获取登录用户信息
func (b *Base) User(c *fiber.Ctx) user.User {
	uid := b._getUID(c)
	if uid != nil {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

// Attempt 尝试登录
func (b *Base) Attempt(c *fiber.Ctx, email string, password string) error {
	// 1. 根据 Email 获取用户
	_user, err := user.GetByEmail(email)

	// 2. 如果出现错误
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在或密码错误")
		} else {
			return errors.New("内部错误，请稍后尝试")
		}
	}

	// 3. 匹配密码
	if !_user.ComparePassword(password) {
		return errors.New("账号不存在或密码错误")
	}

	// 4. 登录用户，保存会话
	//c.Cookie(&fiber.Cookie{
	//	Name:  "uid",
	//	Value: cast.ToString(_user.ID),
	//})
	b.sessionUser(c, _user)
	return nil
}

// Login 登录指定用户
func (b *Base) Login(c *fiber.Ctx, _user user.User) {
	//c.Cookie(&fiber.Cookie{
	//	Name:  "uid",
	//	Value: cast.ToString(_user.ID),
	//})
	b.sessionUser(c, &_user)
}

// Logout 退出用户
func (b *Base) Delete(c *fiber.Ctx) {
	currSession, err := SessionStore.Get(c)
	if err != nil {
		log.Println(err)
	}
	currSession.Delete("uid")
	err = currSession.Save()
	if err != nil {
		log.Println(err)
	}
}

// Check 检测是否登录
func (b *Base) Check(c *fiber.Ctx) bool {
	return b._getUID(c) == nil
}

func (b *Base) sessionUser(c *fiber.Ctx, user *user.User) error {
	currSession, err := SessionStore.Get(c)
	defer currSession.Save()
	if err != nil {
		return err
	}
	err = currSession.Regenerate()
	if err != nil {
		return err
	}
	currSession.Set("uid", user.ID)
	return nil
}
