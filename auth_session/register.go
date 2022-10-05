package auth_session

import (
	"github.com/go-ginseng/ginseng"
	"gorm.io/gorm"
)

const PluginID = "6bdf146d-efa1-4da5-9ace-dce0fdac0c41"

type Option struct {
	DB               *gorm.DB
	SecurityHandlers []SecurityHandler
}

var db *gorm.DB

func RegisterHandler(e *ginseng.Engine, option *Option) {
	db = option.DB
	db.AutoMigrate(&SessionTable{})
	e.Gin().Use(buildSecurityMiddleware(e, option.SecurityHandlers))
}
