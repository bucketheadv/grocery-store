package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return v.(string)
	}
}

func globalPanicHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": errorToString(r),
				})
			}
		}()
		c.Next()
	}
}

func globalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": c.Errors.String(),
			})
			c.Abort()
			return
		}
	}
}

func Load(e *gin.Engine) {
	e.Use(globalPanicHandler())
	e.Use(globalErrorHandler())
}
