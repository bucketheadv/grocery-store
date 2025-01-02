package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/bucketheadv/infragin/components"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Conf struct {
	Server   ServerConf              `json:"server"`
	Apollo   components.ApolloConf   `json:"apollo"`
	XxlJob   components.XxlJobConf   `json:"xxlJob"`
	MySql    MySqlConf               `json:"mysql"`
	Redis    redis.Options           `json:"redis"`
	RocketMQ components.RocketMQConf `json:"rocketMQ"`
}

type ServerConf struct {
	Port int `json:"port"`
}

type MySqlConf struct {
	Url string `json:"url"`
}

type RedisConf struct {
	Url string `json:"url"`
}

var Config Conf

func init() {
	if _, err := toml.DecodeFile("_conf/config.toml", &Config); err != nil {
		logrus.Fatal(err)
	}

	components.InitApolloClient(Config.Apollo, func() {
		var jdbcUrl = components.ApolloApplicationConfig("mysql.url")
		if jdbcUrl != "" {
			Config.MySql.Url = jdbcUrl
		}

		var redisAddr = components.ApolloApplicationConfig("redis.addr")
		if redisAddr != "" {
			Config.Redis.Addr = redisAddr
		}
	})
}
