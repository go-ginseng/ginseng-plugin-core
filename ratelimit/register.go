package ratelimit

import (
	"time"

	"github.com/go-ginseng/ginseng"
)

const PluginID = "b4ec67fa-48d1-408c-a732-2a46638a1558"

type Option struct {
	Period time.Duration
	Limit  int64
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	e.Gin().Use(RateLimit(option.Period, option.Limit))
}
