package utils

import (
	"errors"
	"os"
)

// PathExists 文件目录是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	} else {
		return false, err
	}
}

// CreateFile  创建文件
func CreateFile(filename string) (*os.File, error) {
	exist, err := PathExists(filename)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("file already exist")
	} else {
		return os.Create(filename)
	}
}
