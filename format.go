package main

import (
	"strings"
	"unicode"
)

// keyFilter
func (xj *xjson) keyFilter(str string) (string, string) {
	temp := xj.keyCase(str)
	if _, ok := xj.MapTag[temp]; !ok {
		xj.MapTag[temp] = str
	}
	if xj.JSONTag {
		return str, temp
	}
	xj.MapTag[temp] = ""
	return str, temp

}

// change name to different case
func (xj *xjson) keyCase(str string) string {
	var (
		upperStr string
	)
	temp := strings.FieldsFunc(str, xj.XJSplit)
	for y := 0; y < len(temp); y++ {
		x := []rune(temp[y])
		for i := 0; i < len(x); i++ {
			if i == 0 && y == 0 {
				x[i] = unicode.ToUpper(x[i])
				upperStr += string(x[i])
				continue
			}
			if xj.UpperCase {
				x[i] = unicode.ToUpper(x[i])
				upperStr += string(x[i])
				continue
			}
			if xj.LowerCase {
				x[i] = unicode.ToLower(x[i])
				upperStr += string(x[i])
				continue
			}
			if i == 0 {
				x[i] = unicode.ToUpper(x[i])
				upperStr += string(x[i])
				continue
			}
			upperStr += string(x[i])
			continue
		}
	}
	return upperStr
}

// 4 char
func (xj *xjson) XJSplit(r rune) bool {
	return r == ':' || r == '.' || r == '-' || r == '_'
}
