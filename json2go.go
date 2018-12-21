package main

const (
	// OutTypeForPrint
	OutTypeForPrint = "print"
	// OutTypeForFile
	OutTypeForFile = "file"
)

const (
	// DefaultJsonFile 默认json文件
	DefaultJsonFile = "json2go.json"
	// DefaultOutType 默认输出方式
	DefaultOutType = "print"
	// DefaultOutFile 默认输出文件
	DefaultOutFile = "json2go_types.go"
)

const (
	Xmap    = "map[string]interface {}"
	Xlist   = "[]interface {}"
	Xstring = "string"
	goBegin = "package json2go\n\n"
)

var (
	Xstr     = "%s"
	Xbegin   = "type %s struct {"
	Xend     = "}\n"
	Xkeyv    = "%s %s %s"
	Xkeyvori = "%s %s %s"
	Xkeyvtag = "%s %s `json:\"%s\"`"
)

type xjson struct {
	Name      string
	Msg       string
	LowerCase bool
	UpperCase bool
	JSONTag   bool
	MapTag    map[string]string
	Parent    map[string]interface{}
	Sub       []map[int]string
	Out       []string
	// json文件，默认json2go.json
	JsonFile string
	// 输出类型[print file]，默认print
	OutType string
	// 输出文件，默认json2go_types.go
	OutFile string
}

// New returns a new xjson
func new(name, msg string) *xjson {
	return &xjson{
		Name:    name,
		Msg:     msg,
		JSONTag: true,
		MapTag:  map[string]string{},
	}
}

// reloade
func (xj *xjson) Flush() {
	*xj = *&xjson{
		Name:      xj.Name,
		Msg:       xj.Msg,
		MapTag:    map[string]string{},
		Parent:    map[string]interface{}{},
		Sub:       []map[int]string{},
		Out:       []string{},
		JSONTag:   true,
		LowerCase: false,
		UpperCase: false,
	}
}
