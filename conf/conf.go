package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/bucketheadv/infragin"
	"github.com/bucketheadv/infragin/components"
	"github.com/bucketheadv/infragin/db"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Conf struct {
	Server   infragin.ServerConf     `json:"server"`
	Apollo   components.ApolloConf   `json:"apollo"`
	XxlJob   components.XxlJobConf   `json:"xxlJob"`
	MySql    db.MySqlConf            `json:"mysql"`
	Redis    redis.Options           `json:"redis"`
	RocketMQ components.RocketMQConf `json:"rocketMQ"`
}

var Config Conf

func init() {
	if _, err := toml.DecodeFile("_conf/config.toml", &Config); err != nil {
		logrus.Fatal(err)
	}

	components.InitApolloClient(Config.Apollo, func() {
		components.AssignConfigValueTo("application", "mysql.url", &Config.MySql.Url)
		components.AssignConfigValueTo("application", "redis.url", &Config.Redis.Addr)
	})
}
