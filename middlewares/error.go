package middlewares

import (
	"HereWeGo/common"
	"HereWeGo/initializers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	e := initializers.Engine
	e.Use(globalPanicHandler())
	e.Use(globalErrorHandler())
	e.NoRoute(func(c *gin.Context) {
		var response = common.Response[any]{
			Code:    http.StatusNotFound,
			Message: http.StatusText(http.StatusNotFound),
		}
		common.ApiResponseError(c, response)
	})
}

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
				var response = common.Response[any]{
					Code:    http.StatusInternalServerError,
					Message: errorToString(r),
				}
				common.ApiResponseError(c, response)
			}
		}()
		c.Next()
	}
}

func globalErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			var response = common.Response[any]{
				Code:    http.StatusInternalServerError,
				Message: c.Errors.String(),
			}
			common.ApiResponseError(c, response)
			c.Abort()
			return
		}
	}
}
