package sql_db

import (
	"github.com/go-ginseng/ginseng"
	"gorm.io/gorm"
)

const PluginID = "01f351c9-2908-4b37-9e8f-4f4d7df47138"

var DB *gorm.DB

type Option struct {
	DB *gorm.DB
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	DB = option.DB
}
