package topic

import (
	"fiber-bbs/database"
	"gorm.io/gorm/clause"
)

func GetList(page, pageSize int, maps interface{}, order string) (topics []Topic, count int64, err error) {
	query := database.DB.Preload(clause.Associations).Where(maps)
	switch order {
	case "recent":
		query.Order("created_at desc")
		break
	default:
		query.Order("updated_at desc")
	}
	query.Offset((page - 1) * pageSize).Limit(pageSize).Find(&topics)
	database.DB.Model(&Topic{}).Where(maps).Count(&count)
	return topics, count, nil
}
