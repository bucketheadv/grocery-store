package database

import (
	"github.com/bucketheadv/infra-gin/db"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"grocery-store/conf"
)

var DB *gorm.DB

func init() {
	config, ok := conf.Config.MySQL["main"]
	if !ok {
		logrus.Fatalln("未找到 MySQL: main 配置")
	}
	DB = db.NewMySQL(config, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
