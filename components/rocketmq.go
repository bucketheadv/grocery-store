package components

import (
	"HereWeGo/conf"
	"context"
	"errors"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/sirupsen/logrus"
)

var RocketMqProducer rocketmq.Producer
var RocketMqConsumer rocketmq.PushConsumer

const DemoTopic = "demo_topic"

func init() {
	config := conf.Config.RocketMQ
	if !config.Enabled {
		return
	}
	initProducer()
	initConsumer()
}

func initProducer() {
	config := conf.Config.RocketMQ
	endpoint := config.NameServer
	prod, err := rocketmq.NewProducer(
		producer.WithNameServer(endpoint),
		producer.WithRetry(2),
		producer.WithGroupName("default"),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	err = prod.Start()
	if err != nil {
		logrus.Fatal(err)
	}
	RocketMqProducer = prod
}

func initConsumer() {
	config := conf.Config.RocketMQ
	endpoint := config.NameServer
	consume, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(endpoint),
		consumer.WithRetry(2),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("default"),
	)
	if err != nil {
		logrus.Fatal(err)
	}
	err = consume.Start()
	if err != nil {
		logrus.Fatal(err)
	}
	RocketMqConsumer = consume
}

func SyncSendMsg(msg *primitive.Message) (*primitive.SendResult, error) {
	config := conf.Config.RocketMQ
	if !config.Enabled {
		return nil, errors.New("RocketMQ not enabled")
	}
	return RocketMqProducer.SendSync(context.Background(), msg)
}

func createTopic(topic string) {
	config := conf.Config.RocketMQ
	h, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(config.NameServer)))
	if err != nil {
		logrus.Fatal(err)
	}
	err = h.CreateTopic(context.Background(), admin.WithTopicCreate(topic))
	if err != nil {
		logrus.Println(err)
	}
}

func RegConsumer(topic string, f func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) {
	config := conf.Config.RocketMQ
	if !config.Enabled {
		return
	}
	c := RocketMqConsumer
	err := c.Subscribe(topic, consumer.MessageSelector{}, f)
	if err != nil {
		logrus.Fatal(err)
	}
	err = c.Start()
	if err != nil {
		logrus.Fatal(err)
	}
}
