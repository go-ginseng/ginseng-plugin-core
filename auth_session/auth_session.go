package auth_session

import (
	"github.com/go-ginseng/ginseng"
	"github.com/go-ginseng/ginseng-plugin-core/sql_mem"
)

const PluginID = "6bdf146d-efa1-4da5-9ace-dce0fdac0c41"

type Option struct {
	SecurityHandlers []SecurityHandler
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	e.CheckDependencies(sql_mem.PluginID)
	e.Gin().Use(buildSecurityMiddleware(e, option.SecurityHandlers))
}
