package db

import (
	"HereWeGo/common"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := "root:123456@tcp(localhost:3306)/brian"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalln(err)
	}
}

func Page(db *gorm.DB, page common.BasePage) *gorm.DB {
	return db.Offset(page.Offset()).Limit(page.PageSize)
}
