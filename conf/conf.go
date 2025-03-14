package conf

import (
	"flag"
	"github.com/bucketheadv/infra-core/modules/logger"
	"github.com/bucketheadv/infra-gin/components/apollo"
	"github.com/bucketheadv/infra-gin/conf"
)

var Config conf.Conf

func init() {
	s := flag.String("config", "_conf/config.toml", "配置文件地址")
	flag.Parse()

	if err := conf.Parse(*s, &Config); err != nil {
		logger.Fatal(err)
	}

	apollo.Init(Config.Apollo, func() {
		var mysql = Config.MySQL["main"]
		apollo.AssignApplicationValue("mysql.main.url", &mysql.Url)
		var redis = Config.Redis["main"]
		apollo.AssignApplicationValue("redis.main.url", &redis.Addr)
	})
}
