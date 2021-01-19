package main

import (
	"io/ioutil"
	"log"
	"os"
)

// 生成软连接
func main() {
	err := ioutil.WriteFile("original.txt", []byte("hello world"), 0600)
	if err != nil {
		log.Fatalln(err)
	}

	err = os.Symlink("original.txt", "link.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
