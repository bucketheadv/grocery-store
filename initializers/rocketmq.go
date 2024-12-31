package initializers

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/admin"
	"github.com/apache/rocketmq-client-go/v2/consumer"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"log"
)

var RocketMqProducer rocketmq.Producer
var RocketMqConsumer rocketmq.PushConsumer

var DemoTopic = "demo_topic"

func init() {
	initProducer()
	conf := GetConfig().RocketMQ
	if !conf.Enabled {
		return
	}
	createTopic(DemoTopic)
	initConsumer()
}

func initProducer() {
	conf := GetConfig().RocketMQ
	endpoint := conf.NameServer
	prod, err := rocketmq.NewProducer(
		producer.WithNameServer(endpoint),
		producer.WithRetry(2),
		producer.WithGroupName("default"),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = prod.Start()
	if err != nil {
		log.Fatal(err)
	}
	RocketMqProducer = prod
}

func initConsumer() {
	conf := GetConfig().RocketMQ
	endpoint := conf.NameServer
	consume, err := rocketmq.NewPushConsumer(
		consumer.WithNameServer(endpoint),
		consumer.WithRetry(2),
		consumer.WithConsumerModel(consumer.Clustering),
		consumer.WithGroupName("default"),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = consume.Start()
	if err != nil {
		log.Fatal(err)
	}
	RocketMqConsumer = consume
}

func SyncSendMsg(msg *primitive.Message) (*primitive.SendResult, error) {
	return RocketMqProducer.SendSync(context.Background(), msg)
}

func createTopic(topic string) {
	conf := GetConfig().RocketMQ
	h, err := admin.NewAdmin(admin.WithResolver(primitive.NewPassthroughResolver(conf.NameServer)))
	if err != nil {
		log.Fatal(err)
	}
	err = h.CreateTopic(context.Background(), admin.WithTopicCreate(topic))
	if err != nil {
		log.Println(err)
	}
}

func RegConsumer(topic string, f func(context.Context, ...*primitive.MessageExt) (consumer.ConsumeResult, error)) {
	conf := GetConfig().RocketMQ
	if !conf.Enabled {
		return
	}
	c := RocketMqConsumer
	err := c.Subscribe(topic, consumer.MessageSelector{}, f)
	if err != nil {
		log.Fatal(err)
	}
	err = c.Start()
	if err != nil {
		log.Fatal(err)
	}
}
