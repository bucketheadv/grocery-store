package admin

import (
	"github.com/bucketheadv/infra-gin"
	"github.com/bucketheadv/infra-gin/api"
	"github.com/gin-gonic/gin"
	"grocery-store/model/domain/admin"
	"grocery-store/service/adminsrv"
)

func init() {
	var engine = infra_gin.Engine
	var group = engine.Group("/Admin/Announcement")
	group.GET("/List", func(context *gin.Context) {
		announcement, err := adminsrv.ListAnnouncement()
		if err != nil {
			_ = context.Error(err)
			return
		}
		api.ResponseOk(context, api.Response[[]admin.Announcement]{
			Data: announcement,
		})
	})
}
