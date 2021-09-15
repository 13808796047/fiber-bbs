package upload

import (
	"crypto/md5"
	"errors"
	"fmt"
	"math/rand"
	"mime/multipart"
	"os"
	"path"
	"time"
)

func Save(h *multipart.FileHeader, folder, file_prefix string) (string, error) {

	var err error
	allow_ext := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	ext := path.Ext(h.Filename)
	if _, ok := allow_ext[ext]; !ok {
		errors.New("后缀名不符合上传要求")
		return "", err
	}
	// 创建目录
	uploadDir := fmt.Sprintf("static/upload/%s/%d/%d/%d", folder, time.Now().Year(), time.Now().Month(), time.Now().Day()) + "/"
	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		errors.New("创建文件夹失败")
		return "", err
	}
	//构建文件
	rand.Seed(time.Now().UnixNano())
	randNum := fmt.Sprintf("%d", rand.Intn(9999)+1000)
	hashName := md5.Sum([]byte(time.Now().Format("2006_01_02_15_04_05_") + randNum))
	fileName := file_prefix + "_" + fmt.Sprintf("%x", hashName) + ext
	fpath := uploadDir + fileName
	return fpath, nil
}
