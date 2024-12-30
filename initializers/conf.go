package initializers

import (
	"github.com/BurntSushi/toml"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/go-redis/redis/v8"
	"log"
)

type Conf struct {
	Apollo config.AppConfig `json:"apollo"`
	XxlJob XxlJobConf       `json:"xxlJob"`
	MySql  MySqlConf        `json:"mysql"`
	Redis  redis.Options    `json:"redis"`
}

type XxlJobConf struct {
	Enabled      bool   `json:"enabled"`
	ServerAddr   string `json:"serverAddr"`
	AccessToken  string `json:"accessToken"`
	ExecutorPort string `json:"executorPort"`
	RegistryKey  string `json:"registryKey"`
}

type MySqlConf struct {
	Url string `json:"url"`
}

type RedisConf struct {
	Url string `json:"url"`
}

var conf *Conf

func GetConfig() *Conf {
	if conf == nil {
		if _, err := toml.DecodeFile("_conf/config.toml", &conf); err != nil {
			log.Fatal(err)
		}
	}
	return conf
}
