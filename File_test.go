package File

import (
	"log"
	"testing"
)

func Test_Exists(T *testing.T) {
	var path = "C:/sysgopath/src/github.com/golangframework/File/File.go"
	log.Println(Exists(path))
}
func Test_ReadAllText(T *testing.T) {
	var path = "C:/sysgopath/src/github.com/golangframework/File/File.go"
	log.Println(ReadAllText(path))
}
