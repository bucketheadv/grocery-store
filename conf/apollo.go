package conf

import (
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/apolloconfig/agollo/v4/storage"
	"github.com/sirupsen/logrus"
)

type ApolloChangeListener struct{}

func (c *ApolloChangeListener) OnChange(event *storage.ChangeEvent) {
	for k, v := range event.Changes {
		logrus.Infof("apollo %v config changed, key: %v, old value: %v, new value: %v",
			event.Namespace, k, v.OldValue, v.NewValue)
	}
}

func (c *ApolloChangeListener) OnNewestChange(event *storage.FullChangeEvent) {
	for k, v := range event.Changes {
		logrus.Infof("apollo config pull, key: %s, value: %v", k, v)
	}
}

var ApolloClient agollo.Client

func InitApolloClient() {
	c := Config.Apollo

	client, err := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		var appConfig = &config.AppConfig{
			AppID:          c.AppID,
			Cluster:        c.Cluster,
			NamespaceName:  c.NamespaceName,
			IP:             c.IP,
			IsBackupConfig: c.IsBackupConfig,
		}
		return appConfig, nil
	})

	if err != nil {
		logrus.Infof("初始化Apollo失败, %s", err.Error())
		return
	}

	client.AddChangeListener(&ApolloChangeListener{})
	ApolloClient = client
}

func GetApolloConfig(key string) string {
	return ApolloClient.GetValue(key)
}
