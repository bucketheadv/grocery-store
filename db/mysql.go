package db

import (
	"HereWeGo/conf"
	"HereWeGo/core"
	"database/sql"
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

func Page(db *gorm.DB, page core.Page) *gorm.DB {
	return db.Offset(page.Offset()).Limit(page.Limit())
}

func CloseRows(rows *sql.Rows) {
	err := rows.Close()
	if err != nil {
		logrus.Println(err)
	}
}
