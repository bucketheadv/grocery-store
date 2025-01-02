package consumer

import (
	"HereWeGo/initializer"
	"context"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/sirupsen/logrus"
)

func init() {
	topic := initializer.DemoTopic
	client := initializer.RocketMQConsumer
	client.RegConsumer(topic, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range ext {
			logrus.Infof("消费到topic: %s, ext: %s", topic, ext[i])
		}
		return consumer.ConsumeSuccess, nil
	})
}
