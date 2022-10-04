package static

import (
	_static "github.com/gin-contrib/static"
	"github.com/go-ginseng/ginseng"
)

const PluginID = "29dae50c-ad68-4a16-90de-487fa5942678"

type StaticConfig struct {
	RoutePrefix string
	FilePath    string
}

type Option struct {
	StaticConfigs []StaticConfig

	// if yes, any route without /api/ will append .html to the request route
	// your api route should have /api/ prefix
	RootHTML bool
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	for _, config := range option.StaticConfigs {
		e.AppendMiddleware(_static.Serve(config.RoutePrefix, _static.LocalFile(config.FilePath, false)))
	}
	if option.RootHTML {
		e.AppendMiddleware(htmlMiddleware)
	}
}
