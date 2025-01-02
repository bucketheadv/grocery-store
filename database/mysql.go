package database

import (
	"HereWeGo/conf"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func init() {
	config := conf.Config.MySql
	var err error
	DB, err = gorm.Open(mysql.Open(config.Url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		logrus.Fatalln(err)
	}
}
