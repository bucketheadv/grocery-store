package database

import (
	"github.com/bucketheadv/infra-gin/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"grocery-store/conf"
)

var DB *gorm.DB

func init() {
	config := conf.Config.MySQL["main"]
	DB = db.NewMySQL(*config, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
