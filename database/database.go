package database

import (
	"fiber-bbs/config"
	"fmt"
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

var DB *gorm.DB
func ConnectionDB() *gorm.DB {
	var err error
	// 初始化MySQL连接信息
	var (
		host = config.Config("DB_HOST")
		port = config.Config("DB_PORT")
		database = config.Config("DB_DATABASE")
		username = config.Config("DB_USERNAME")
		password = config.Config("DB_PASSWORD")
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%t&loc=%s",
		username, password, host, port, database,"utf8mb4", true, "Local")
	gormConfig := mysql.New(mysql.Config{
		DSN: dsn,
	})
	var level gormlogger.LogLevel
	if cast.ToBool(config.Config("APP_DEBUG")) {
		// 读取不到数据也会显示
		level = gormlogger.Warn
	} else {
		// 只有错误才会显示
		level = gormlogger.Error
	}

	// 准备数据库连接池
	DB, err = gorm.Open(gormConfig, &gorm.Config{
		Logger: gormlogger.Default.LogMode(level),
	})
	if err != nil {
		fmt.Println("数据库连接错误:",err)
	}
	return DB
}