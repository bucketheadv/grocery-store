package service

import (
	"fmt"
	"github.com/bucketheadv/infra-core/modules/logger"
	"github.com/bucketheadv/infra-gin"
	"github.com/bucketheadv/infra-gin/db"
	"gorm.io/gorm"
	"grocery-store/database"
	"grocery-store/model/po"
	"slices"
	"time"
)

const (
	userCacheKey     = "user:cache:%d"
	userPageCacheKey = "user:page:cache:%d:%d"
)

func GetUser(id int) (po.User, error) {
	var key = fmt.Sprintf(userCacheKey, id)
	data, err := db.FetchCache(database.RedisClient, key, 5*time.Minute, func() (po.User, error) {
		var user po.User
		rows, err := database.DB.Where("id = ?", id).Find(&user).Rows()
		if err != nil {
			return user, err
		}
		if !rows.Next() {
			return user, gorm.ErrRecordNotFound
		}
		defer db.CloseRows(rows)
		return user, nil
	})
	return data, err
}

func GetUsers(ids []int) ([]po.User, error) {
	if len(ids) == 0 {
		return make([]po.User, 0), nil
	}

	var result = make([]po.User, 0)
	var missingIds = make([]int, 0)
	var keys = make([]string, 0)
	for _, id := range ids {
		var key = fmt.Sprintf(userCacheKey, id)
		keys = append(keys, key)
	}
	foundUsers, err := db.GetCaches[po.User](database.RedisClient, keys)
	if err != nil {
		panic(err)
	}
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
		var users []po.User
		rows, err := database.DB.Where("id in (?)", missingIds).Find(&users).Rows()
		if err != nil {
			return nil, err
		}
		defer db.CloseRows(rows)
		for _, user := range users {
			var key = fmt.Sprintf(userCacheKey, user.ID)
			err := db.SetCache(database.RedisClient, key, user, 5*time.Minute)
			if err != nil {
				panic(err)
			}
			result = append(result, user)
		}
	}
	return result, nil
}

func UserByPage(page infra_gin.Page) (infra_gin.PageResult[po.User], error) {
	var key = fmt.Sprintf(userPageCacheKey, page.PageNo, page.PageSize)
	data, err := db.FetchCache(database.RedisClient, key, 5*time.Minute, func() (*[]po.User, error) {
		var users *[]po.User
		rows, err := db.Page(database.DB, page).Find(&users).Rows()
		if err != nil {
			logger.Error("查询数据失败, ", err.Error())
			return nil, err
		}
		defer db.CloseRows(rows)
		return users, nil
	})
	if err != nil {
		panic(err)
	}
	pageInfo := infra_gin.PageResult[po.User]{
		Page:    page,
		Records: *data,
	}
	return pageInfo, nil
}
