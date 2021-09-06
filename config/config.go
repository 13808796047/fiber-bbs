package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("读取配置文件出错:",err)
	}
	return os.Getenv(key)
}