package main

import (
	"fmt"
	"reflect"
)

// subStruct
func (xj *xjson) subStruct(key string, out interface{}) {
	dout := out.(map[string]interface{})
	m := map[int]string{}
	index := 0
	sk := xj.keyCase(key)
	m[index] = fmt.Sprintf(Xbegin, sk)
	for k, v := range dout {
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
		index++
		m[index] = fmt.Sprintf(Xkeyv, k, tmptype, xj.MapTag[k])
	}
	m[index+1] = Xend
	xj.Sub = append(xj.Sub, m)
}

// subList
func (xj *xjson) subList(key string, out interface{}) {
	dout := out.([]interface{})
	list := map[string]interface{}{}
	for _, vl := range dout {
		vtype := fmt.Sprintf(Xstr, reflect.TypeOf(vl))
		if vtype == Xlist {
			xj.subList(key, vl)
			continue
		}
		if vtype == Xmap {
			m := vl.(map[string]interface{})
			for k, v := range m {
				_, k := xj.keyFilter(k)
				if _, ok := list[k]; ok {
					continue
				}
				list[k] = v
			}
		}
	}
	xj.subStruct(key, list)
}
