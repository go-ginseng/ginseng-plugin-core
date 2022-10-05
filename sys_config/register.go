package sys_config

import (
	"github.com/go-ginseng/ginseng"
	"gorm.io/gorm"
)

const PluginID = "27eb1542-ce62-4e6c-b284-fbda8c6bca2d"

var db *gorm.DB
var mem *gorm.DB

type Option struct {
	DB  *gorm.DB
	MEM *gorm.DB
}

func RegisterHandler(e *ginseng.Engine, option *Option) {
	db = option.DB
	mem = option.MEM
	db.AutoMigrate(&SystemConfig{})
	mem.AutoMigrate(&SystemConfig{})

	e.AppendPreRunFunc(SyncMem)
}
