package api

import (
	"HereWeGo/components"
	"HereWeGo/conf"
	"HereWeGo/core"
	"HereWeGo/db/model"
	"HereWeGo/service"
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"math/rand"
	"strconv"
	"strings"
)

func init() {
	r := components.Engine
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
		core.ApiResponseOk(c, core.Response[*model.User]{
			Data: user,
		})
	})

	group.GET("/Query", func(c *gin.Context) {
		page := core.ParsePageParams(c)
		pageInfo, err := service.UserByPage(page)
		if err != nil {
			_ = c.Error(errors.New("查询用户失败, " + err.Error()))
			return
		}
		core.ApiResponseOk(c, core.Response[core.PageResult[model.User]]{
			Data: pageInfo,
		})
	})

	group.GET("/QueryByIds", func(c *gin.Context) {
		ids := strings.Split(c.Query("ids"), ",")
		idsInt := make([]int, len(ids))
		for i, id := range ids {
			idsInt[i], _ = strconv.Atoi(id)
		}
		users, _ := service.GetUsers(idsInt)
		core.ApiResponseOk(c, core.Response[[]model.User]{
			Data: users,
		})
	})

	group.GET("/Apollo", func(c *gin.Context) {
		var timeout = conf.ApolloNamespaceValue[int]("application", "timeout")
		core.ApiResponseOk(c, core.Response[int]{
			Data: timeout,
		})
	})

	group.GET("/SendMqMsg", func(c *gin.Context) {
		var msg = fmt.Sprintf("测试数据 %d", rand.Int())
		_, err := components.SyncSendMsg(&primitive.Message{
			Topic: components.DemoTopic,
			Body:  []byte(msg),
		})
		if err != nil {
			logrus.Error(err)
		}
		core.ApiResponseOk(c, core.Response[*model.User]{
			Data: nil,
		})
	})
}
