package files

import (
	"cron/config"
	util "cron/pkg/security"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func LogFile() (io.Writer, string) {
	// 目录路径
	logDirPath := config.Conf.App.LogDir
	// 文件名
	logFileName := config.Conf.App.AppName
	// 文件路径
	logFilePath := path.Join(logDirPath, logFileName)
	//创建文件夹·
	if ok, _ := FileExists(logFilePath); !ok {
		if err := CreateSavePath(logDirPath, 0755); err != nil {
			fmt.Println("create log file path failed")
		}
	}
	// 获取文件句柄
	f, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("Open Src File err", err)
	}
	return f, logFilePath
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	if os.IsNotExist(err) {
		return false, err
	}
	return true, err
}

type FileType int

const TypeImage FileType = iota + 1

func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

func GetFileExt(name string) string {
	return path.Ext(name)
}

func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsNotExist(err)
}

func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return err
}

func Basename(path string) string {
	return filepath.Base(path)
}
