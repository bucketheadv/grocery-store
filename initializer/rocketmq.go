package initializer

import (
	"HereWeGo/conf"
	"github.com/bucketheadv/infragin/components"
)

const DemoTopic = "demo_topic"

var RocketMQProducer components.InfraRocketMQProducer
var RocketMQConsumer components.InfraRocketMQConsumer

func init() {
	config := conf.Config.RocketMQ
	RocketMQProducer = components.InitProducer(config)
	RocketMQConsumer = components.InitConsumer(config)
}
