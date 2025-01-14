package service

import (
	"fmt"
	"github.com/bucketheadv/infra-gin"
	"github.com/bucketheadv/infra-gin/db"
	"github.com/sirupsen/logrus"
	"grocery-store/database"
	"grocery-store/database/model"
	"slices"
	"time"
)

const (
	userCacheKey     = "user:cache:%d"
	userPageCacheKey = "user:page:cache:%d:%d"
)

func GetUser(id int) (*model.User, error) {
	var key = fmt.Sprintf(userCacheKey, id)
	var data = db.FetchCache(database.RedisClient, key, 5*time.Minute, func() model.User {
		var user model.User
		rows, err := database.DB.Where("id = ?", id).Find(&user).Rows()
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
	var foundUsers = db.GetCaches[model.User](database.RedisClient, keys)
	var foundUserIds = make([]int, 0)
	for _, u := range foundUsers {
		foundUserIds = append(foundUserIds, u.ID)
		result = append(result, u)
	}
	for _, id := range ids {
		if !slices.Contains(foundUserIds, id) {
			missingIds = append(missingIds, id)
		}
	}

	if len(missingIds) > 0 {
		var users []model.User
		rows, err := database.DB.Where("id in (?)", missingIds).Find(&users).Rows()
		if err != nil {
			return nil, err
		}
		defer db.CloseRows(rows)
		for _, user := range users {
			var key = fmt.Sprintf(userCacheKey, user.ID)
			db.SetCache(database.RedisClient, key, user, 5*time.Minute)
			result = append(result, user)
		}
	}
	return result, nil
}

func UserByPage(page infra_gin.Page) (infra_gin.PageResult[model.User], error) {
	var key = fmt.Sprintf(userPageCacheKey, page.PageNo, page.PageSize)
	var data = db.FetchCache(database.RedisClient, key, 5*time.Minute, func() *[]model.User {
		var users *[]model.User
		rows, err := db.Page(database.DB, page).Find(&users).Rows()
		if err != nil {
			logrus.Error("查询数据失败, ", err.Error())
			return nil
		}
		defer db.CloseRows(rows)
		return users
	})
	pageInfo := infra_gin.PageResult[model.User]{
		Page:    page,
		Records: *data,
	}
	return pageInfo, nil
}
