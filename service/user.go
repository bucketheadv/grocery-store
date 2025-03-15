package service

import (
	"fmt"
	"github.com/bucketheadv/infra-gin"
	"github.com/bucketheadv/infra-gin/db"
	"gorm.io/gorm"
	"grocery-store/database"
	"grocery-store/model/po"
	"time"
)

const (
	userCacheKey     = "user:cache:%d"
	userPageCacheKey = "user:page:cache:%d:%d"
)

func GetUser(id int) (po.User, error) {
	var key = fmt.Sprintf(userCacheKey, id)
	data, err := db.FetchCache(database.RedisClient, key, 1*time.Minute, func() (po.User, error) {
		var user po.User
		database.DB.Where("id = ?", id).Find(&user)
		return user, nil
	})
	return data, err
}

func GetUsers(ids []int) ([]po.User, error) {
	return db.GetModelCaches[po.User](database.RedisClient, userCacheKey, ids, 1*time.Minute, func(missingIds []int) *gorm.DB {
		return database.DB.Where("id IN ?", missingIds)
	})
}

func UserByPage(page infra_gin.Page) (infra_gin.PageResult[po.User], error) {
	var key = fmt.Sprintf(userPageCacheKey, page.PageNo, page.PageSize)
	return db.FetchCache(database.RedisClient, key, 5*time.Minute, func() (infra_gin.PageResult[po.User], error) {
		return db.Page[po.User](database.DB, page)
	})
}
