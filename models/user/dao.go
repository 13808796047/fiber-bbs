package user

import (
	"fiber-bbs/database"
	"fiber-bbs/pkgs/password"
	"github.com/spf13/cast"
)

func (u *User) Create() (err error) {
	if err = database.DB.Create(&u).Error; err != nil {
		return err
	}
	return
}
func Get(idstr interface{}) (User, error) {
	var user User
	id := cast.ToInt(idstr)

	if err := database.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}
func GetByEmail(idstr string) (*User, error) {
	var user *User
	if err := database.DB.Where("email=?", idstr).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

// ComparePassword 对比密码是否匹配
func (u *User) ComparePassword(_password string) bool {
	return password.CheckHash(_password, u.Password)
}
