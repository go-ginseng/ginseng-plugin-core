package sql_db

import (
	"github.com/go-ginseng/ginseng"
	"gorm.io/gorm"
)

const PluginID = "39f8218d-345c-44fa-9c14-329371be9386"

var DB *gorm.DB

type Option struct {
	DB *gorm.DB
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	DB = option.DB
}
