# oozlog-oozg
[![Build Status](https://travis-ci.org/usth/oozlog.svg?branch=master)](https://travis-ci.org/usthooz/oozlog)
[![Go Report Card](https://goreportcard.com/badge/github.com/usthooz/oozlog)](https://goreportcard.com/report/github.com/usthooz/oozlog)
[![GoDoc](http://godoc.org/github.com/usthooz/oozlog?status.svg)](http://godoc.org/github.com/usthooz/oozlog/go)

Golang的简洁日志工具。

## install
```
$ go get -u github.com/usthooz/oozlog
```

## 功能
1. 支持日志级别设置  
2. 多级别输出  
3. 输出到文件(待续)

## 使用
```
package main

import(
    uoozg "github.com/usthooz/oozlog/go"
)

funca main(){
    // 设置日志级别
    uoozg.Default.Level = uoozg.DebugLevel
	uoozg.Infof("infof test->%d", uoozg.DebugLevel)
	uoozg.Debugf("debug test->%d", uoozg.DebugLevel)
	uoozg.Warnf("warnf test->%d", uoozg.DebugLevel)
	uoozg.Errorf("error test->%d", uoozg.DebugLevel)
	uoozg.Fatalf("fatalf test->%d", uoozg.DebugLevel)
}
```  
  
效果图:
![uzthoozlogg](https://github.com/usthooz/oozlog/blob/master/img/golang.png)
