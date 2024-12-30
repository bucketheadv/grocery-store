package controller

import (
	"HereWeGo/common"
	"HereWeGo/service"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

func UserController(r *gin.Engine) {
	group := r.Group("/user")
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

	group.GET("/query", func(c *gin.Context) {
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
}
