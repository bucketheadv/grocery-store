package initial

import (
	"github.com/bucketheadv/infra-gin/middlewares"
	"github.com/gin-gonic/gin"
	"grocery-store/filter"
)

var Engine *gin.Engine

func init() {
	Engine = gin.Default()
	middlewares.RegErrorHandler(Engine)
	Engine.Use(filter.LoginFilter())
}
