package admin

import (
	"github.com/bucketheadv/infra-gin/api"
	"github.com/gin-gonic/gin"
	"grocery-store/constants/errcode"
	"grocery-store/initial"
	"grocery-store/model/params"
	"grocery-store/service/adminsrv"
)

func init() {
	var engine = initial.Engine
	group := engine.Group("/Admin/Auth")
	group.GET("/Login", func(context *gin.Context) {
		var p params.LoginParam
		err := context.BindJSON(&p)
		if err != nil {
			_ = context.Error(api.NewParamError("参数错误"))
			return
		}
		token, err := adminsrv.AuthUser(p)
		if err != nil {
			_ = context.Error(errcode.ErrInvalidNameOrPassword)
			return
		}
		api.ResponseOk(context, api.Response[string]{
			Data: token,
		})
	})
}
