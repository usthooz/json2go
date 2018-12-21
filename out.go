package main

import (
	"fmt"
	"os"

	sh "github.com/codeskyblue/go-sh"
	"github.com/usthooz/oozlog/go"
)

func (xj *xjson) appendStr(kv string) {
	xj.Out = append(xj.Out, kv)
}

// printStruct
func (xj *xjson) printStruct() {
	for _, value := range xj.Sub {
		for i := 0; i < len(value); i++ {
			fmt.Println(value[i])
		}
	}
	for i := 0; i < len(xj.Out); i++ {
		fmt.Println(xj.Out[i])
	}
}

// write struct to file
func (xj *xjson) writefileStruct() {
	if len(xj.OutFile) == 0 {
		xj.OutFile = DefaultOutFile
	}
	file, err := os.OpenFile(xj.OutFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	file.WriteString(goBegin)
	if err != nil {
		ozlog.Errorf("new file error...")
	}
	for _, value := range xj.Sub {
		for i := 0; i < len(value); i++ {
			if i == 0 || i == (len(value)-1) {
				file.WriteString(value[i] + " " + "\n")
			} else {
				file.WriteString("	" + value[i] + " " + "\n")
			}
		}
	}

	for i := 0; i < len(xj.Out); i++ {
		if i == 0 || i == (len(xj.Out)-1) {
			file.WriteString(xj.Out[i] + "\n")
		} else {
			file.WriteString("	" + xj.Out[i] + "\n")
		}
	}
	// 格式化
	sh.Command("gofmt", "-w", ".", sh.Dir(xj.OutFile)).Run()
}
