package service

import (
	"HereWeGo/common"
	"HereWeGo/db"
	"HereWeGo/db/model"
	"fmt"
	"log"
)

const (
	userCacheKey     = "user:cache:%d"
	userPageCacheKey = "user:page:cache:%d:%d"
)

func GetUser(id int) (*model.User, error) {
	var key = fmt.Sprintf(userCacheKey, id)
	var data = db.CacheByKey(key, func() model.User {
		var user model.User
		rows, err := db.DB.Where("id = ?", id).Find(&user).Rows()
		if err != nil {
			panic(err)
		}
		defer db.CloseRows(rows)
		return user
	})
	return &data, nil
}

func UserByPage(page common.Page) (*common.PageResult[model.User], error) {
	var key = fmt.Sprintf(userPageCacheKey, page.PageNo, page.PageSize)
	var data = db.CacheByKey(key, func() *[]model.User {
		var users *[]model.User
		rows, err := db.Page(db.DB, page).Find(&users).Rows()
		if err != nil {
			log.Println("查询数据失败, ", err.Error())
			return nil
		}
		defer db.CloseRows(rows)
		return users
	})
	pageInfo := &common.PageResult[model.User]{
		Page:    page,
		Records: *data,
	}
	return pageInfo, nil
}
