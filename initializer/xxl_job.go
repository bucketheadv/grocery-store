package initializer

import (
	"github.com/bucketheadv/infra-gin/components/xxljob"
	"grocery-store/conf"
)

var XxlJobClient xxljob.Client

func init() {
	config := conf.Config.XxlJob
	client := xxljob.NewClient(config)
	XxlJobClient = client
}
