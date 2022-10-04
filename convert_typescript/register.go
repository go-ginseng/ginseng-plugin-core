package convert_typescript

import (
	"os"

	"github.com/go-ginseng/ginseng"
)

const PluginID = "5ae74879-576b-4bf6-a505-a311b589dc23"

type Option struct {
	OutDir string
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	registeredTypes = make(map[string][]interface{})
	e.AppendPreRunFunc(func() {
		os.MkdirAll(option.OutDir, os.ModePerm)
		Generate(option.OutDir)
	})
}
