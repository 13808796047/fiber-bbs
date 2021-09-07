package auth

import (
	"errors"
	"fiber-bbs/models/user"

	// "fiber-bbs/pkgs/session"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"

	"gorm.io/gorm"
)

func _getUID(c *fiber.Ctx) string {

	_uid := c.Cookies("uid")

	if len(_uid) > 0 {
		return _uid
	}
	return ""
}

// 获取登录用户信息
func User(c *fiber.Ctx) user.User {
	uid := _getUID(c)
	if len(uid) > 0 {
		_user, err := user.Get(uid)
		if err == nil {
			return _user
		}
	}
	return user.User{}
}

// Attempt 尝试登录
func Attempt(c *fiber.Ctx, email string, password string) error {
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
	c.Cookie(&fiber.Cookie{
		Name:  "uid",
		Value: cast.ToString(_user.ID),
	})
	return nil
}

// Login 登录指定用户
func Login(c *fiber.Ctx, _user user.User) {
	c.Cookie(&fiber.Cookie{
		Name:  "uid",
		Value: cast.ToString(_user.ID),
	})

}

// Logout 退出用户
func Logout(c *fiber.Ctx) {
	c.ClearCookie("uid")
}

// Check 检测是否登录
func Check(c *fiber.Ctx) bool {
	return len(_getUID(c)) > 0
}
