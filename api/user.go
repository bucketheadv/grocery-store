package api

import (
	"fmt"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/bucketheadv/infra-core/basic"
	"github.com/bucketheadv/infra-core/modules/logger"
	"github.com/bucketheadv/infra-gin/api"
	"github.com/bucketheadv/infra-gin/components/apollo"
	"github.com/gin-gonic/gin"
	"grocery-store/initial"
	"grocery-store/model/domain"
	"grocery-store/service"
	"math/rand"
	"net/http"
	"strings"
)

func init() {
	r := initial.Engine
	group := r.Group("/User")
	group.GET("/GetById", func(c *gin.Context) {
		id, err := api.GetQuery[int](c, "id")
		if err != nil {
			_ = c.Error(err)
			return
		}
		user, err := service.GetUser(id)
		if err != nil {
			_ = c.Error(api.NewBizError(http.StatusInternalServerError, "查询数据失败, "+err.Error()))
			return
		}
		api.ResponseOk(c, api.Response[domain.User]{
			Data: user,
		})
	})

	group.GET("/Query", func(c *gin.Context) {
		page := api.ParsePageParams(c)
		pageInfo, err := service.UserByPage(page)
		if err != nil {
			_ = c.Error(api.NewBizError(http.StatusInternalServerError, "查询用户失败, "+err.Error()))
			return
		}
		api.ResponseOk(c, api.Response[api.PageResult[domain.User]]{
			Data: pageInfo,
		})
	})

	group.GET("/QueryByIds", func(c *gin.Context) {
		var p = strings.Split(c.Query("id"), ",")
		ids, err := basic.ArrayElemTo[int64](p)
		if err != nil {
			_ = c.Error(err)
			return
		}
		users, _ := service.GetUsers(ids)
		api.ResponseOk(c, api.Response[[]domain.User]{
			Data: users,
		})
	})

	group.GET("/Apollo", func(c *gin.Context) {
		var timeout = apollo.NamespaceValue[int]("application", "timeout")
		api.ResponseOk(c, api.Response[int]{
			Data: timeout,
		})
	})

	group.GET("/SendMqMsg", func(c *gin.Context) {
		var msg = fmt.Sprintf("测试数据 %d", rand.Int())
		_, err := initial.RocketMQProducer.SendSync(&primitive.Message{
			Topic: initial.DemoTopic,
			Body:  []byte(msg),
		})
		if err != nil {
			logger.Error(err)
		}
		api.ResponseOk(c, api.Response[any]{})
	})
}
