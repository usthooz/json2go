package main

import (
	"fmt"
	"os"
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
func (xj *xjson) writefileStruct(fileName string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	file.WriteString(goBegin)
	if err != nil {
		fmt.Println("new file error...")
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
}
