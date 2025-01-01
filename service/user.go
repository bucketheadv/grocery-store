package service

import (
	"HereWeGo/common"
	"HereWeGo/db"
	"HereWeGo/db/model"
	"fmt"
	"log"
	"slices"
	"time"
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

func GetUsers(ids []int) ([]model.User, error) {
	if len(ids) == 0 {
		return make([]model.User, 0), nil
	}

	var result = make([]model.User, 0)
	var missingIds = make([]int, 0)
	var keys = make([]string, 0)
	for _, id := range ids {
		var key = fmt.Sprintf(userCacheKey, id)
		keys = append(keys, key)
	}
	var foundUsers = db.GetCaches[model.User](keys)
	var foundUserIds = make([]int, 0)
	for _, u := range foundUsers {
		foundUserIds = append(foundUserIds, u.Id)
		result = append(result, u)
	}
	for _, id := range ids {
		if !slices.Contains(foundUserIds, id) {
			missingIds = append(missingIds, id)
		}
	}

	if len(missingIds) > 0 {
		var users []model.User
		rows, err := db.DB.Where("id in (?)", missingIds).Find(&users).Rows()
		if err != nil {
			return nil, err
		}
		defer db.CloseRows(rows)
		for _, user := range users {
			var key = fmt.Sprintf(userCacheKey, user.Id)
			db.SetCache(key, user, 5*time.Minute)
			result = append(result, user)
		}
	}
	return result, nil
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
