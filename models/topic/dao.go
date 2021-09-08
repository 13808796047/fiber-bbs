package topic

import (
	"fiber-bbs/database"
	"github.com/spf13/cast"
	"gorm.io/gorm/clause"
)

func GetList(data map[string]interface{}) (topics *[]Topic, count int64, err error) {
	database.DB.Preload(clause.Associations).Offset(cast.ToInt(data["page"])).Limit(cast.ToInt(data["per_page"])).Find(&topics)
	database.DB.Model(&Topic{}).Count(&count)
	return topics, count, nil
}

//func GetCount() int64{
//
//}
