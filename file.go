package File

import (
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

func ReadAllText(filename string) string {
	buff, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("文件读取失败")
	}
	return string(buff)
}
