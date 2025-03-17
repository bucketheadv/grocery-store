package service

import (
	"fmt"
	"github.com/bucketheadv/infra-gin/api"
	"github.com/bucketheadv/infra-gin/db"
	"gorm.io/gorm"
	"grocery-store/database"
	"grocery-store/model/domain"
	"time"
)

const (
	userCacheKey     = "user:cache:%d"
	userPageCacheKey = "user:page:cache:%d:%d"
)

func GetUser(id int) (domain.User, error) {
	var key = fmt.Sprintf(userCacheKey, id)
	data, err := db.FetchCache(database.RedisClient, key, 1*time.Minute, func() (domain.User, error) {
		var user domain.User
		var err = database.DB.Where("id = ?", id).Find(&user).Error
		return user, err
	})
	return data, err
}

func GetUsers(ids []int64) ([]domain.User, error) {
	return db.ModelCaches[domain.User](database.RedisClient, userCacheKey, ids, 1*time.Minute, func(missingIds []int64) *gorm.DB {
		return database.DB.Where("id IN ?", missingIds)
	})
}

func UserByPage(page api.Page) (api.PageResult[domain.User], error) {
	var key = fmt.Sprintf(userPageCacheKey, page.PageNo, page.PageSize)
	return db.FetchCache(database.RedisClient, key, 5*time.Minute, func() (api.PageResult[domain.User], error) {
		return db.Page[domain.User](database.DB, page)
	})
}
