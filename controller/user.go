package controller

import (
	"HereWeGo/common"
	"HereWeGo/initializers"
	"HereWeGo/service"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func init() {
	r := initializers.Engine
	group := r.Group("/User")
	group.GET("/GetById", func(c *gin.Context) {
		id, success := c.GetQuery("id")
		if !success {
			_ = c.Error(errors.New("参数错误"))
			return
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			_ = c.Error(errors.New("参数转换错误"))
			return
		}
		user, err := service.GetUser(idInt)
		if err != nil {
			_ = c.Error(errors.New("查询数据失败, " + err.Error()))
			return
		}
		common.ApiResponseOK(c, gin.H{
			"data": user,
		})
	})

	group.GET("/Query", func(c *gin.Context) {
		page := common.ParsePageParams(c)
		pageInfo, err := service.UserByPage(page)
		if err != nil {
			_ = c.Error(errors.New("查询用户失败, " + err.Error()))
			return
		}
		common.ApiResponseOK(c, gin.H{
			"data": pageInfo,
		})
	})

	group.GET("/Apollo", func(c *gin.Context) {
		conf := initializers.ApolloClient.GetConfig("application")
		var timeout = conf.GetIntValue("timeout", 0)
		common.ApiResponseOK(c, gin.H{
			"data": timeout,
		})
	})
}
