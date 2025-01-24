package conf

import (
	"github.com/bucketheadv/infra-gin/components/apollo"
	"github.com/bucketheadv/infra-gin/conf"
	"github.com/sirupsen/logrus"
)

var Config conf.Conf

func init() {
	if err := conf.Parse("_conf/config.toml", &Config); err != nil {
		logrus.Fatal(err)
	}

	apollo.Init(Config.Apollo, func() {
		var mysql = Config.MySQL["main"]
		apollo.AssignApplicationValue("mysql.main.url", &mysql.Url)
		var redis = Config.Redis["main"]
		apollo.AssignApplicationValue("redis.main.url", &redis.Addr)
	})
}
