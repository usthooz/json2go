package main

import (
	"fmt"
	"os"
)

// readJsonToFile 从文件中读取json信息
func readJsonToFile(fileName string) string {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("open file err")
	}
	var (
		json string
	)
	buf := make([]byte, 1024)
	for {
		len, _ := file.Read(buf)
		if len == 0 {
			break
		}
		json = `` + string(buf[:len]) + ``
	}
	return json
}
