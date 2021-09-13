package category

import (
	"fiber-bbs/database"
	"gorm.io/gorm"
)

func Get(id int) (*Category, error) {
	var category Category
	err := database.DB.Model(&Category{}).Where("id=?", id).First(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &category, nil
}
func List() ([]*Category, error) {
	category := []*Category{}
	err := database.DB.Model(&Category{}).Find(&category).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return category, nil
}
