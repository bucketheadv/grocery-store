package database

import (
	"github.com/bucketheadv/infra-core/modules/logger"
	"github.com/bucketheadv/infra-gin/db"
	"gorm.io/gorm"
	log "gorm.io/gorm/logger"
	"grocery-store/conf"
)

var DB *gorm.DB

func init() {
	config, ok := conf.Config.MySQL["main"]
	if !ok {
		logger.Fatal("未找到 MySQL: main 配置")
	}
	DB = db.NewMySQL(config, &gorm.Config{
		Logger: log.Default.LogMode(log.Info),
	})
}
