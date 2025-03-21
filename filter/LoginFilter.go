package filter

import (
	"github.com/bucketheadv/infra-gin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

var whiteList = []string{
	"/Admin/Auth/Login",
}

func LoginFilter() gin.HandlerFunc {
	return func(c *gin.Context) {
		for _, p := range whiteList {
			if p == c.Request.URL.Path {
				break
			}
			var auth = c.GetHeader("Authorization")
			if auth == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, api.Response[string]{
					Code:    http.StatusUnauthorized,
					Message: "未登录",
				})
			}
		}
		c.Next()
	}
}
