package bootstrap

import (
	"fiber-bbs/config"
	"fiber-bbs/database"
	"fiber-bbs/models/category"
	"fiber-bbs/models/user"
	"fiber-bbs/routes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func SetupRoute(app *fiber.App) {
	routes.RegisterWebRoutes(app)
}

// SetupDB 初始化数据库和ORM
func SetupDB() {
	// 建立数据库连接池
	db := database.ConnectionDB()
	// 命令行打印数据库请求的信息
	sqlDB, _ := db.DB()
	// 设置最大连接数
	sqlDB.SetMaxOpenConns(cast.ToInt(config.Config("DB_MAX_IDLE_CONNECTIONS")))
	// 设置最大空闲连接数
	sqlDB.SetMaxIdleConns(cast.ToInt(config.Config("DB_MAX_OPEN_CONNECTIONS")))
	// 设置每个链接的过期时间
	sqlDB.SetConnMaxLifetime(time.Duration(cast.ToInt(config.Config("DB_MAX_LIFE_SECONDS"))) * time.Second)
	// 创建和维护数据表结构
	migration(db)
}
func migration(db *gorm.DB) {
	//var categories = []category.Category{{
	//	Name:        "分享",
	//	Description: "分享创造，分享发现",
	//},
	//	{
	//		Name:        "教程",
	//		Description: "开发技巧、推荐扩展包等",
	//	},
	//	{
	//		Name:        "问答",
	//		Description: "请保持友善，互帮互助",
	//	},
	//	{
	//		Name:        "公告",
	//		Description: "站点公告",
	//	}}
	//db.Model(&category.Category{}).Create(&categories)
	// 自动迁移
	db.AutoMigrate(&user.User{}, &category.Category{})
}
