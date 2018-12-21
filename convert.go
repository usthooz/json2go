package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	ozlog "github.com/usthooz/oozlog/go"
)

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
	if len(xj.OutType) == 0 {
		xj.OutType = DefaultOutType
	}
	if xj.OutType == OutTypeForPrint {
		xj.printStruct()
	} else if xj.OutType == OutTypeForFile {
		if len(xj.OutFile) == 0 {
			xj.OutFile = DefaultOutFile
		}
		xj.writefileStruct(xj)
	}
	xj.Flush()
	return xj
}