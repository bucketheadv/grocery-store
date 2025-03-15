package admin

import (
	"github.com/bucketheadv/infra-gin"
	"github.com/bucketheadv/infra-gin/api"
	"github.com/gin-gonic/gin"
)

func init() {
	var engine = infra_gin.Engine
	var group = engine.Group("/Admin/Announcement")
	group.GET("/List", func(context *gin.Context) {
		api.ResponseOk(context, api.Response[string]{
			Data: "OK",
		})
	})
}
