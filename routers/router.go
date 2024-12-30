package routers

import (
	"HereWeGo/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	e := initializers.Engine
	InitRouter(e)
}

func InitRouter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404 Not Found",
		})
	})
}
