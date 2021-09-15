package topic

import (
	"fiber-bbs/database"
	"gorm.io/gorm/clause"
	"log"
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
func Get(id string) (topic *Topic, err error) {
	if err := database.DB.Preload(clause.Associations).First(&topic, id).Error; err != nil {
		log.Fatal(err)
	}

	return
}
func Delete(id string) error {
	var err error
	if err = database.DB.Where("id=?", id).Delete(&Topic{}).Error; err == nil {
		return nil
	}
	return err
}
func Update(maps interface{}, topic *Topic) (err error) {
	if err = database.DB.Where(maps).Save(&topic).Error; err == nil {
		return nil
	}
	return
}
