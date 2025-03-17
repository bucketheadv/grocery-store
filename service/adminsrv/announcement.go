package adminsrv

import (
	"github.com/bucketheadv/infra-gin/db"
	"grocery-store/database"
	"grocery-store/model/domain/admin"
	"time"
)

const announcementCacheKey = "announcement_cache"

var dbs = database.DB

func ListAnnouncement() ([]admin.Announcement, error) {
	return db.FetchCache(database.RedisClient, announcementCacheKey, 5*time.Minute, func() ([]admin.Announcement, error) {
		var tx = dbs.Where("status = ?", 1).Order("id DESC")
		var result []admin.Announcement
		var err = tx.Find(result).Error
		return result, err
	})
}
