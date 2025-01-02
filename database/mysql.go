package database

import (
	"HereWeGo/conf"
	"github.com/bucketheadv/infragin/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	config := conf.Config.MySql
	DB = db.NewMySQL(config, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
}
