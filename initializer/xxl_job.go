package initializer

import (
	"HereWeGo/conf"
	"github.com/bucketheadv/infragin/components/xxljob"
)

var XxlJobClient xxljob.Client

func init() {
	config := conf.Config.XxlJob
	client := xxljob.NewClient(config)
	XxlJobClient = client
}
