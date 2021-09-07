package password

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

// 密码加密
func Hash(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal("密码加密错误:", err)
	}
	return string(bytes)
}

// 对比明文密码和数据库的哈希值
func CheckHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Fatal("密码对比错误", err)
	}
	return err == nil
}

// 判断字符串是否是哈希过的数据
func IsHashed(str string) bool {
	return len(str) == 60
}
