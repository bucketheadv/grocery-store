package initializer

import (
	"HereWeGo/conf"
	"github.com/bucketheadv/infragin/components/rocket"
)

const DemoTopic = "demo_topic"

var RocketMQProducer rocket.InfraRocketMQProducer
var RocketMQConsumer rocket.InfraRocketMQConsumer

func init() {
	config := conf.Config.RocketMQ
	RocketMQProducer = rocket.InitProducer(config)
	RocketMQConsumer = rocket.InitConsumer(config)
}
