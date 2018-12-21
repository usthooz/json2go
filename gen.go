package main

import (
	"os"

	ozlog "github.com/usthooz/oozlog/go"
)

// readJsonAndGen 从文件中读取json信息并且生成struct
func readJsonAndGen(jsonFile, outType, outFile string) {
	if len(jsonFile) == 0 {
		jsonFile = DefaultJsonFile
	}
	file, err := os.OpenFile(jsonFile, os.O_RDONLY, 0666)
	if err != nil {
		ozlog.Errorf("open file err")
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
		ck := New(json, jsonFile, outType, outFile)
		ck.json2Struct()
	}
	return
}
