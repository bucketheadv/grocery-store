package db

import (
	"HereWeGo/common"
	"HereWeGo/initializers"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func init() {
	conf := initializers.GetConfig().MySql
	var err error
	DB, err = gorm.Open(mysql.Open(conf.Url), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func Page(db *gorm.DB, page common.Page) *gorm.DB {
	return db.Offset(page.Offset()).Limit(page.PageSize)
}
