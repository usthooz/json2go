package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	"github.com/usthooz/oozlog/go"
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
		ck := New("json2go", json)
		ck.json2Struct()
	}
	return json
}

// json2Struct 完成json转换
func (xj *xjson) json2Struct() *xjson {
	Xkeyv = Xkeyvori
	if xj.JSONTag {
		Xkeyv = Xkeyvtag
	}
	if err := json.Unmarshal([]byte(xj.Msg), &xj.Parent); err != nil {
		ozlog.Errorf("parse json failed:", err)
		return nil
	}
	xj.appendStr(fmt.Sprintf(Xbegin, xj.Name))
	for k, v := range xj.Parent {
		orik, k := xj.keyFilter(k)
		tmptype := fmt.Sprintf(Xstr, reflect.TypeOf(v))
		if tmptype == Xmap {
			xj.subStruct(orik, v)
			tmptype = k
		}
		if tmptype == Xlist {
			xj.subList(orik, v)
			tmptype = k
		}
		xj.appendStr(fmt.Sprintf(Xkeyv, k, tmptype, xj.MapTag[k]))
	}
	xj.printStruct()
	xj.Flush()
	return xj
}
