package initializers

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"log"
)

type ApolloChangeListener struct{}

func (c *ApolloChangeListener) OnChange(event *storage.ChangeEvent) {
	for k, v := range event.Changes {
		log.Printf("apollo %v config changed, key: %v, old value: %v, new value: %v\n",
			event.Namespace, k, v.OldValue, v.NewValue)
	}
}

func (c *ApolloChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	for k, v := range event.Changes {
		log.Printf("apollo config pull, key: %s, value: %v\n", k, v)
	}
}

var ApolloClient agollo.Client

func init() {
	conf := &config.AppConfig{
		AppID:          "SampleApp",
		Cluster:        "DEV",
		IP:             "http://localhost:8080",
		NamespaceName:  "application",
		IsBackupConfig: true,
	}

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return conf, nil
	})

	if err != nil {
		log.Printf("初始化Apollo失败, %s\n", err.Error())
		return
	}

	client.AddChangeListener(&ApolloChangeListener{})
	ApolloClient = client

	log.Println("初始化Apollo成功")
}
