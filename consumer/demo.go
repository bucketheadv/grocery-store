package consumer

import (
	"HereWeGo/initializers"
	"context"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/sirupsen/logrus"
)

func init() {
	topic := initializers.DemoTopic
	initializers.RegConsumer(topic, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range ext {
			logrus.Debugf("消费到topic: %s, ext: %s", topic, ext[i])
		}
		return consumer.ConsumeSuccess, nil
	})
}
