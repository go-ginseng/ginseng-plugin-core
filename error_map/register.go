package error_map

import "github.com/go-ginseng/ginseng"

const PluginID = "14c3bd37-2bc4-4312-be7e-5ddd137e3ce0"

const (
	MODE_INT = "int"
	MODE_STR = "str"
)

var MODE = ""

var intMap map[int]string
var strMap map[string]string

type Option struct {
	Mode string
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	MODE = option.Mode
	switch MODE {
	case MODE_INT:
		intMap = make(map[int]string)
	case MODE_STR:
		strMap = make(map[string]string)
	default:
		panic("error_map: invalid mode")
	}
}
