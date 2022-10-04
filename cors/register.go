package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/go-ginseng/ginseng"
)

const PluginID = "52910d5b-6214-4596-8b5a-8a3a3af94855"

type Option struct {
	AllowAllOrigins        bool
	AllowOrigins           []string
	AllowOriginFunc        func(origin string) bool
	AllowMethods           []string
	AllowHeaders           []string
	AllowCredentials       bool
	ExposeHeaders          []string
	MaxAge                 time.Duration
	AllowWildcard          bool
	AllowBrowserExtensions bool
	AllowWebSockets        bool
	AllowFiles             bool
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	e.Gin().Use(cors.New(cors.Config{
		AllowAllOrigins:        option.AllowAllOrigins,
		AllowOrigins:           option.AllowOrigins,
		AllowOriginFunc:        option.AllowOriginFunc,
		AllowMethods:           option.AllowMethods,
		AllowHeaders:           option.AllowHeaders,
		AllowCredentials:       option.AllowCredentials,
		ExposeHeaders:          option.ExposeHeaders,
		MaxAge:                 option.MaxAge,
		AllowWildcard:          option.AllowWildcard,
		AllowBrowserExtensions: option.AllowBrowserExtensions,
		AllowWebSockets:        option.AllowWebSockets,
		AllowFiles:             option.AllowFiles,
	}))
}
