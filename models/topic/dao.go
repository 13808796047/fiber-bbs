package topic

import (
	"fiber-bbs/database"
	"gorm.io/gorm/clause"
)

func GetList(page, pageSize int, maps interface{}) (topics *[]Topic, count int64, err error) {
	database.DB.Preload(clause.Associations).Where(maps).Offset(page).Limit(pageSize).Find(&topics)
	database.DB.Model(&Topic{}).Where(maps).Count(&count)
	return topics, count, nil
}

//func GetByCategory(page,pageSize int,maps interface{}) (topics *[]Topic, count int64, err error) {
//	page := data["page"]
//	database.DB.Preload(clause.Associations).Offset(page).Limit(data["per_page"])).Find(&topics)
//	database.DB.Model(&Topic{}).Count(&count)
//	return topics, count, nil
//}
