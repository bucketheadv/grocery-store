package api

import (
	"errors"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/bucketheadv/infra-core/basic"
	"github.com/bucketheadv/infra-core/modules/logger"
	"github.com/bucketheadv/infra-gin"
	"github.com/bucketheadv/infra-gin/components/apollo"
	"github.com/gin-gonic/gin"
	"grocery-store/initializer"
	"grocery-store/model/po"
	"grocery-store/service"
	"math/rand"
	"strings"
)

func init() {
	r := infra_gin.Engine
	group := r.Group("/User")
	group.GET("/GetById", func(c *gin.Context) {
		id, err := infra_gin.GetQuery[int](c, "id")
		if err != nil {
			_ = c.Error(err)
			return
		}
		user, err := service.GetUser(id)
		if err != nil {
			_ = c.Error(errors.New("查询数据失败, " + err.Error()))
			return
		}
		infra_gin.ApiResponseOk(c, infra_gin.Response[po.User]{
			Data: user,
		})
	})

	group.GET("/Query", func(c *gin.Context) {
		page := infra_gin.ParsePageParams(c)
		pageInfo, err := service.UserByPage(page)
		if err != nil {
			_ = c.Error(errors.New("查询用户失败, " + err.Error()))
			return
		}
		infra_gin.ApiResponseOk(c, infra_gin.Response[infra_gin.PageResult[po.User]]{
			Data: pageInfo,
		})
	})

	group.GET("/QueryByIds", func(c *gin.Context) {
		var p = strings.Split(c.Query("id"), ",")
		ids, err := basic.ArrayElemTo[int](p)
		if err != nil {
			_ = c.Error(err)
			return
		}
		users, _ := service.GetUsers(ids)
		infra_gin.ApiResponseOk(c, infra_gin.Response[[]po.User]{
			Data: users,
		})
	})

	group.GET("/Apollo", func(c *gin.Context) {
		var timeout = apollo.NamespaceValue[int]("application", "timeout")
		infra_gin.ApiResponseOk(c, infra_gin.Response[int]{
			Data: timeout,
		})
	})

	group.GET("/SendMqMsg", func(c *gin.Context) {
		var msg = fmt.Sprintf("测试数据 %d", rand.Int())
		_, err := initializer.RocketMQProducer.SendSync(&primitive.Message{
			Topic: initializer.DemoTopic,
			Body:  []byte(msg),
		})
		if err != nil {
			logger.Error(err)
		}
		infra_gin.ApiResponseOk(c, infra_gin.Response[any]{})
	})
}
