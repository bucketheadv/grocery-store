package initializer

import (
	"github.com/bucketheadv/infra-core/modules/logger"
	"github.com/bucketheadv/infra-gin/components/rocket"
	"grocery-store/conf"
)

const DemoTopic = "demo_topic"

var RocketMQProducer rocket.InfraRocketMQProducer
var RocketMQConsumer rocket.InfraRocketMQConsumer

func init() {
	config, ok := conf.Config.RocketMQ["main"]
	if !ok {
		logger.Fatal("未找到 RocketMQ: main 配置")
	}
	RocketMQProducer = rocket.InitProducer(config)
	RocketMQConsumer = rocket.InitConsumer(config)
}
