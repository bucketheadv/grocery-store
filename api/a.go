package api

import (
	"github.com/bucketheadv/infra-gin"
	"github.com/bucketheadv/infra-gin/middlewares"
)

func init() {
	// a.go 保证能优于该包路径下其他文件先加载，本组件必须先加载才能生效
	middlewares.RegErrorHandler(infra_gin.Engine)
}
