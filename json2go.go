package main

const (
	Xmap    = "map[string]interface {}"
	Xlist   = "[]interface {}"
	Xstring = "string"
	goBegin = "package json2struct\n\n"
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
}

// New returns a new xjson
func New(name, msg string) *xjson {
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
