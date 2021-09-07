package models

import "fiber-bbs/database"

func (u *User) Create() (err error) {
	if err = database.DB.Create(&u).Error; err != nil {
		return err
	}
	return
}
