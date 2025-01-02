package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/bucketheadv/infragin"
	"github.com/bucketheadv/infragin/components/apollo"
	"github.com/bucketheadv/infragin/components/rocket"
	"github.com/bucketheadv/infragin/components/xxljob"
	"github.com/bucketheadv/infragin/db"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

type Conf struct {
	Server   infragin.ServerConf `json:"server"`
	Apollo   apollo.Conf         `json:"apollo"`
	XxlJob   xxljob.Conf         `json:"xxlJob"`
	MySql    db.MySqlConf        `json:"mysql"`
	Redis    redis.Options       `json:"redis"`
	RocketMQ rocket.Conf         `json:"rocketMQ"`
}

var Config Conf

func init() {
	if _, err := toml.DecodeFile("_conf/config.toml", &Config); err != nil {
		logrus.Fatal(err)
	}

	apollo.InitClient(Config.Apollo, func() {
		apollo.AssignConfigValueTo("application", "mysql.url", &Config.MySql.Url)
		apollo.AssignConfigValueTo("application", "redis.url", &Config.Redis.Addr)
	})
}
