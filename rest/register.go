package rest

import (
	"github.com/go-ginseng/ginseng"
	"github.com/go-ginseng/ginseng-plugin-core/error_map"
)

const PluginID = "83e01b1c-34c3-463a-a276-ead8ff40e0b6"

type Option struct{}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	e.CheckDependencies(error_map.PluginID)
}
