package file

import (
	"io/ioutil"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
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
	CheckErr(err)
	return string(buff)
}
