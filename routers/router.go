package routers

import (
	"HereWeGo/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter(r *gin.Engine) {
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "404 Not Found",
		})
	})

	controller.UserController(r)
}
