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
	DefaultOutFile = "gen_json2go_types.go"
	// DefaultStructName 结构体名称
	DefaultStructName = "Json2GoAutoGenerate"
)

const (
	Xmap    = "map[string]interface {}"
	Xlist   = "[]interface {}"
	Xstring = "string"
	goBegin = "package json2go\n\n"
)

var (
	Xstr     = "%s"
	Xbegin   = "\ntype %s struct {"
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
func New(msg, jsonFile, outType, outFile string) *xjson {
	return &xjson{
		Name:     DefaultStructName,
		Msg:      msg,
		JSONTag:  true,
		MapTag:   map[string]string{},
		JsonFile: jsonFile,
		OutType:  outType,
		OutFile:  outFile,
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
