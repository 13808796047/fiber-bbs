package models

import (
	"fiber-bbs/pkgs/password"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	*gorm.Model
	Name     string `gorm:"type:varchar(255);not null;unique" valid:"name"`
	Email    string `gorm:"type:varchar(255);unique;" valid:"email"`
	Password string `gorm:"type:varchar(255)" valid:"password"`

	// gorm:"-" —— 设置 GORM 在读写时略过此字段，仅用于表单验证
	PasswordConfirm string `gorm:"-" valid:"password_confirm" form:"password_confirm"`
	Capatcha        string `gorm:"-" valid:"captcha" form:"captcha"`
}

// BeforeSave Gorm的模型钩子,在保存和更新模型前调用
func (u *User) BeforeSave(tx *gorm.DB) error {
	if !password.IsHashed(u.Password) {
		u.Password = password.Hash(u.Password)
	}
	return nil
}
