package consumer

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/bucketheadv/infra-core/modules/logger"
	"grocery-store/initializer"
)

func init() {
	topic := initializer.DemoTopic
	client := initializer.RocketMQConsumer
	client.RegConsumer(topic, func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
		for i := range ext {
			logger.Infof("消费到topic: %s, ext: %s", topic, ext[i])
		}
		return consumer.ConsumeSuccess, nil
	})
}
