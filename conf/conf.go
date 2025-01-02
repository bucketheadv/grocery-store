package conf

import (
	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis/v8"
	"log"
)

type Conf struct {
	Server   ServerConf    `json:"server"`
	Apollo   ApolloConf    `json:"apollo"`
	XxlJob   XxlJobConf    `json:"xxlJob"`
	MySql    MySqlConf     `json:"mysql"`
	Redis    redis.Options `json:"redis"`
	RocketMQ RocketMQConf  `json:"rocketMQ"`
}

type ServerConf struct {
	Port int `json:"port"`
}

type ApolloConf struct {
	Enabled        bool   `json:"enabled"`
	AppID          string `json:"appId"`
	Cluster        string `json:"cluster"`
	NamespaceName  string `json:"namespaceName"`
	IP             string `json:"ip"`
	IsBackupConfig bool   `default:"true" json:"isBackupConfig"`
}

type XxlJobConf struct {
	Enabled      bool   `json:"enabled"`
	ServerAddr   string `json:"serverAddr"`
	AccessToken  string `json:"accessToken"`
	ExecutorPort string `json:"executorPort"`
	RegistryKey  string `json:"registryKey"`
	LogDir       string `json:"logDir"`
	LogRetention int    `json:"logRetention"`
}

type MySqlConf struct {
	Url string `json:"url"`
}

type RedisConf struct {
	Url string `json:"url"`
}

type RocketMQConf struct {
	Enabled    bool     `json:"enabled"`
	NameServer []string `json:"nameServer"`
}

var Config Conf

func init() {
	if _, err := toml.DecodeFile("_conf/config.toml", &Config); err != nil {
		log.Fatal(err)
	}

	initApolloClient(Config.Apollo, func() {
		var jdbcUrl = ApolloApplicationConfig("mysql.url")
		if jdbcUrl != "" {
			Config.MySql.Url = jdbcUrl
		}

		var redisAddr = ApolloApplicationConfig("redis.addr")
		if redisAddr != "" {
			Config.Redis.Addr = redisAddr
		}
	})
}
