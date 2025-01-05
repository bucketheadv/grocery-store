package conf

import (
	"github.com/bucketheadv/infragin/components/apollo"
	"github.com/bucketheadv/infragin/conf"
	"github.com/sirupsen/logrus"
)

var Config conf.Conf

func init() {
	if err := conf.Parse("_conf/config.toml", &Config); err != nil {
		logrus.Fatal(err)
	}

	apollo.InitClient(Config.Apollo, func() {
		var mysql = Config.MySql["main"]
		apollo.AssignConfigValueTo("application", "mysql.main.url", &mysql.Url)
		var redis = Config.Redis["main"]
		apollo.AssignConfigValueTo("application", "redis.main.url", &redis.Addr)
	})
}
