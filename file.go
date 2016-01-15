package File

import (
	"errors"
	"io/ioutil"
	"os"
)

func Exists(filename string) bool {
	_, err := os.Open(filename)
	if err != nil && os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
func ReadAllBytes(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("文件读取失败")
	}
	return ioutil.ReadAll(f)
}
func ReadAllText(filename string) (string, error) {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.New("文件读取失败")
	}
	return string(buff), nil
}
