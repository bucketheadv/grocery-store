package initializer

import (
	"github.com/bucketheadv/infra-gin/components/rocket"
	"grocery-store/conf"
)

const DemoTopic = "demo_topic"

var RocketMQProducer rocket.InfraRocketMQProducer
var RocketMQConsumer rocket.InfraRocketMQConsumer

func init() {
	config := conf.Config.RocketMQ["main"]
	RocketMQProducer = rocket.InitProducer(*config)
	RocketMQConsumer = rocket.InitConsumer(*config)
}
